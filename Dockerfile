FROM golang:latest
RUN mkdir -p /home/the-devops-project/ 
WORKDIR /home/the-devops-project
COPY . .
RUN go build -o main .
EXPOSE 8082
CMD ["./main"]