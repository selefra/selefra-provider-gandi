# Table: gandi_certificate

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cn | string | X | √ | Common Name. | 
| package_name | string | X | √ |  | 
| status | string | X | √ | One of: 'pending', 'valid', 'revoked', 'replaced', 'replaced_rev', 'expired'. | 
| expiration | timestamp | X | √ | Expiration date. | 
| id | string | X | √ | UUID. | 


