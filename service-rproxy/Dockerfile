FROM nginx:1.15.11 as production-stage
RUN chown nginx:nginx /usr/share/nginx/html
EXPOSE 80 443
CMD ["nginx", "-g", "daemon off;"]