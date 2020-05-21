package search

import (
	"context"
	"encoding/json"
	"log"
	"os"

	Log "github.com/jampajeen/go-cqrs-taxi/logger"
	"github.com/jampajeen/go-cqrs-taxi/schema"
	"github.com/olivere/elastic"
)

const mapping = `{
	"settings":{
		"number_of_shards":1,
		"number_of_replicas":0
	},
	"mappings":{
		"taxi":{
			"properties":{
				"body":{
					"type":"keyword"
				},
				"lat":{
					"type":"double"
				},
				"lon":{
					"type":"double"
				},
				"location":{
					"type":"geo_point"
				}
			}
		}
	}
}`

type ElasticRepository struct {
	client *elastic.Client
}

func NewElastic(url string) (*ElasticRepository, error) {
	ctx := context.Background()
	client, err := elastic.NewClient(
		elastic.SetTraceLog(log.New(os.Stdout, "", 0)),
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}

	exists, err := client.IndexExists("taxies").Do(ctx)
	if err != nil {
		Log.Fatal(err)
	}
	if exists {
		_, err = client.DeleteIndex("taxies").Do(ctx)
		if err != nil {
			Log.Fatal(err)
		}
	}
	_, err = client.CreateIndex("taxies").BodyString(mapping).Do(ctx)
	if err != nil {
		Log.Fatal(err)
	}

	return &ElasticRepository{client}, nil
}

func (r *ElasticRepository) Close() {
}

func (r *ElasticRepository) UpdateTaxiLocation(ctx context.Context, id string, lat float64, lon float64) error {
	Log.Debug("ID=%s, Lat=%s, Lon=%s", id, lat, lon)
	_, err := r.client.Update().Index("taxies").Type("taxi").Id(id).
		Script(elastic.NewScript("ctx._source.location.lat = params.lat; ctx._source.location.lon = params.lon").Params(map[string]interface{}{"lat": lat, "lon": lon})).
		Do(ctx)
	return err
}

func (r *ElasticRepository) InsertTaxi(ctx context.Context, taxi schema.Taxi) error {
	Log.Debug("taxi=%+v", taxi)
	_, err := r.client.Index().
		Index("taxies").
		Type("taxi").
		Id(taxi.ID).
		BodyJson(taxi).
		Refresh("wait_for").
		Do(ctx)
	return err
}

func (r *ElasticRepository) SearchTaxies(ctx context.Context, query string, start uint64, size uint64) ([]schema.Taxi, error) {
	point, err := elastic.GeoPointFromString(query)
	if err != nil {
		Log.Fatal(err)
		return nil, err
	}

	gdq := elastic.NewGeoDistanceQuery("location").
		GeoPoint(point).
		Distance("50km")

	bq := elastic.NewBoolQuery().
		Must(
			elastic.NewMatchAllQuery(),
		).
		Filter(
			gdq,
		)

	result, err := r.client.
		Search().
		Index("taxies").
		Type("taxi").
		Query(bq).
		From(int(start)).
		Size(int(size)).
		Do(ctx)

	if err != nil {
		Log.Fatal(err)
	}

	if err != nil {
		return nil, err
	}

	taxies := []schema.Taxi{}
	for _, hit := range result.Hits.Hits {
		var taxi schema.Taxi
		if err = json.Unmarshal(*hit.Source, &taxi); err != nil {
			Log.Error(err)
		}
		taxies = append(taxies, taxi)
	}
	return taxies, nil
}

func (r *ElasticRepository) SearchTaxiesByKeyword(ctx context.Context, query string, start uint64, size uint64) ([]schema.Taxi, error) {
	result, err := r.client.Search().
		Index("taxies").
		Query(
			elastic.NewMultiMatchQuery(query, "body").
				Fuzziness("3").
				PrefixLength(1).
				CutoffFrequency(0.0001),
		).
		From(int(start)).
		Size(int(size)).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	taxies := []schema.Taxi{}
	for _, hit := range result.Hits.Hits {
		var taxi schema.Taxi
		if err = json.Unmarshal(*hit.Source, &taxi); err != nil {
			Log.Error(err)
		}
		taxies = append(taxies, taxi)
	}
	return taxies, nil
}
