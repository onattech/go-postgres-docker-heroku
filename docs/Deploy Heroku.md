### Log in to Container Registry:
```cmd
heroku container:login
```

### Build the image and push to Container Registry:
```cmd
heroku container:push web -a <heroku-app-name>
```

### Then release the image to your app:
```cmd
heroku container:release web -a <heroku-app-name>
```



