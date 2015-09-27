# AUTH

## Creating Users
```
curl -XPOST localhost:8080/api/registration -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE0NDMyNzQyODUsImlzcyI6Ind3dy5hc3ZpbnMuY29tLmJyIiwic2NvcGUiOiJhZG1pbiIsInN1YiI6InZpbmljaXVzQGFzdmlucy5jb20uYnIifQ.LeovKEWm806y-t0oWTYV9QVmynahlRM50Hw3Bcg-MHI' -H 'Content-Type: application/json' -d '{"email": "paciente@example.com", "first_name": "Jose", "last_name": "Silva", "scope": "patient", "password": "asasdaads"}'
```

## Loggin in

```
curl -XPOST localhost:8080/api/login -H 'Content-Type: application/json' -d '{"email":"paciente@example.com", "password": "asasdaads"}'
```

## Checking if token is valid

```
curl -H 'Content-Type: application/json' localhost:8080/api/userinfo -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE0NDMzODg1NjksImlzcyI6Ind3dy5hc3ZpbnMuY29tLmJyIiwic2NvcGUiOiJwYXRpZW50Iiwic3ViIjoicGFjaWVudGVAZXhhbXBsZS5jb20ifQ.88aAzXPNua4NVZO8B1RWODFyYMF6sC787NwgPe3M9P4'
```
