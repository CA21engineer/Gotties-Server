package firebase


//色々と初期化しまくってる
config := &firebase.Config{
StorageBucket: "<BUCKET_NAME>.appspot.com",
}

opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
app, err := firebase.NewApp(context.Background(), config, opt)
if err != nil {
log.Fatalln(err)
}

client, err := app.Storage(context.Background())
if err != nil {
log.Fatalln(err)
}

bucket, err := client.DefaultBucket()
if err != nil {
log.Fatalln(err)
}

//=======================
//firebase sdk git
//https://github.com/firebase/firebase-admin-go

//firebase sdk doc
//https://firebase.google.com/docs/admin/setup/



export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/service-account-file.json"

app, err := firebase.NewApp(context.Background(), nil)
if err != nil {
log.Fatalf("error initializing app: %v\n", err)
}


// Initialize default app
app, err := firebase.NewApp(context.Background(), nil)
if err != nil {
log.Fatalf("error initializing app: %v\n", err)
}

// Access auth service from the default app
client, err := app.Auth(context.Background())
if err != nil {
log.Fatalf("error getting Auth client: %v\n", err)
}