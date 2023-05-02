## Copy env

Copy the env.example to .env and fill the required key

---
## Commands
1. Go modules
    ```
    go mode tidy
    ```

2. Migration
    - Create new migration file
      ```
      go run migration/create.go create_user_table
      ```

    - Run migration up
      ```
      migrate -database postgres://user:password@host/db_name -path schema up
      ```

    - Run migration down (all)
      ```
      migrate -database postgres://user:password@host/db_name -path schema down
      ```

    - Run migration down (specific)
      ```
      migrate -database postgres://user:password@host/db_name -path schema goto <version>
      ```

3. Run
    ```
    go run main.go
    ```

4. Mock repository for unit test
    ```
    mockery --name=RepositoryInterfaceName
    ```