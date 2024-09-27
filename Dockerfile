FROM debian:stable-slim

COPY personal-website /bin/personal-website

#COPY /templates /templates

ENV PORT=8080

CMD ["/bin/personal-website"]