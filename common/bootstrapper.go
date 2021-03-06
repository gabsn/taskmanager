package common

func StartUp() {
    // Initialize AppConfig variable
    initConfig()
    // Initialize RSA keys for JWT Authentication
    initKeys()
    // Start a MongoDB session
    createDbSession()
    // Add indexes into MongoDB
    addIndexes()
}
