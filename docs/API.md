
### Request
```json
{
    "email": "me@hyunsang.dev",
    "passowrd": "q1w2e3r4!",
    "phone_number": "010-1234-1234",
    "birthday": "2004-06-25",
    "name": "박현상",
    "display_name": "토모",
    "AllowMarketing": false
}
```

### Response
```json
{
    "stats": {
        "code": 200,
        "stats": "ok",
        "message": "성공적으로 새로운 사용자의 정보를 만들었습니다."
    },
    "data": {
        "Pk": 0,
        "Id": "20f862f9-5fcd-4f1a-9b2a-b5938853ea75",
        "Name": "박현상",
        "Email": "me@hyunsang.dev",
        "Password": "$2a$10$KPV2PhTi3AUo9kTJt76soe.9OReg/undu1ZbrvCIRyZoLbAPe/Ksm",
        "PhoneNumber": "010-1234-1234",
        "DisplayName": "토모",
        "Birthday": "2004-06-25T00:00:00Z",
        "AllowMarketing": {
            "Bool": false,
            "Valid": false
        },
        "CreatedAt": "2024-01-12T16:59:22.027+09:00",
        "UpdatedAt": "2024-01-12T16:59:22.027+09:00",
        "ActivatedAt": {
            "Time": "0001-01-01T00:00:00Z",
            "Valid": false
        },
        "DeletedAt": {
            "Time": "0001-01-01T00:00:00Z",
            "Valid": false
        }
    },
    "responded_at": "2024-01-12T16:59:22.033024+09:00"
}
```