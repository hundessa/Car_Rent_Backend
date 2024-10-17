package firebase

import (
    "context"
    "firebase.google.com/go/v4"
    "firebase.google.com/go/v4/storage"
    "google.golang.org/api/option"
    "log"
)

func InitFirebase() (*storage.Client, error) {
    // Path to your service account key JSON file
    opt := option.WithCredentialsFile("/path/to/your/serviceAccountKey.json")

    // Initialize the Firebase app
    app, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
        log.Fatalf("error initializing Firebase app: %v", err)
        return nil, err
    }

    // Initialize Firebase Storage client
    storageClient, err := app.Storage(context.Background())
    if err != nil {
        log.Fatalf("error initializing Firebase Storage: %v", err)
        return nil, err
    }

    return storageClient, nil
}
