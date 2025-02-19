# stage de build
FROM public.ecr.aws/docker/library/golang:alpine3.19 AS build

WORKDIR /app

COPY . /app

#RUN GOOS=linux go build -o todoapi main.go

# stage imagem final
FROM scratch

WORKDIR /app

COPY --from=build /app/todoapi ./

EXPOSE 8081

CMD [ "./todoapi" ]
