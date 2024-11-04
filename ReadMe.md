# GoTik

GoTik is an app for getting video from tiktok.

## Backend

1. Prepare the environment variables

    ```yaml
    PORT=8080
    NATS_URL=nats://127.0.0.1:422
    ```

2. NATS is required, since the app will send the job to queue
3. Run the test

    ```shell
    go test ./apps/... --cover
    ```

4. Run the app

    ```shell
    go run server.go
    ```

## Frontend

1. Prepare the environment variables

    ```yaml
    VITE_API_URL=http://localhost:8080/query
    VITE_API_URL_WS=ws://localhost:8080/query
    ```

2. Run the test

    ```shell
    yarn test:unit
    ```

3. Run the app

    ```shell
    yarn dev
    ```

Check the video sample: <https://github.com/ariefsn/gotik/raw/refs/heads/main/GoTik.mp4>
