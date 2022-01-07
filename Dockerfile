# Sttart from base image : 1.16.12
FROM golang:1.16.12

ENV ELASTIC_HOSTS=localhost:9200
ENV LOG_LEVEL=info

# Configure the repo url so we can configure our work directory:
ENV REPO_URL=github.com/Delaram-Gholampoor-Sagha/bookstore_items_api

# Setup out $GOPATH
ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/$REPO_URL


# Copy the entire source code from the current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src



COPY . $WORKPATH

WORKDIR $WORKPATH

RUN go build -o items-api .

# Expose port 8081 to the world:
EXPOSE 8081

CMD ["./items-api"]