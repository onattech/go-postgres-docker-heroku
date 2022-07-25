### Log in to Container Registry:
```cmd
heroku container:login
```

### Build the image and push to Container Registry:
```cmd
heroku container:push web -a go-fiber-gorm-postgresql
```

### Then release the image to your app:
```cmd
heroku container:release web -a go-fiber-gorm-postgresql
```



