# Table: gandi_domain

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ | Unique id of the domain. | 
| tld | string | X | √ | The top-level domain. | 
| owner | string | X | √ | The username of the owner. | 
| dates_created_at | timestamp | X | √ | The date the domain started to be handled by Gandi. | 
| livednssec_available | bool | X | √ | Indicates if DNSSEC with liveDNS may be applied to this domain. | 
| dates_pending_delete_ends_at | timestamp | X | √ | The date from which the domain will be available after a deletion. | 
| dates_registry_ends_at | timestamp | X | √ | The date the domain will end at the registry. | 
| fqdn | string | X | √ | Fully qualified domain name, written in its native alphabet (IDN). | 
| domain_owner | string | X | √ | The full name of the owner. | 
| auto_renew | bool | X | √ | Automatic renewal status. | 
| sharing_id | string | X | √ | The id of the organization. | 
| status | json | X | √ | The status of the domain. | 
| dates_deletes_at | timestamp | X | √ | The date on which the domain will be deleted at the registry. | 
| fqdn_unicode | string | X | √ | Fully qualified domain name, written in unicode. | 
| dates_hold_begins_at | timestamp | X | √ | The date from which the domain is held. | 
| livedns_current | string | X | √ | Type of nameservers currently set. classic corresponds to Gandi's classic nameservers, livedns is for the new, default, Gandi nameservers, premium_dns indicates the presence of Gandi's Premium DNS nameserver and the corresponding service subscription, and other is for custom nameservers. | 
| dnssec_available | bool | X | √ | Indicates if DNSSEC may be applied to the domain. | 
| orga_owner | string | X | √ | The username of the organization owner. | 
| tags | json | X | √ | Tags associated to this domain. | 
| dates_registry_created_at | timestamp | X | √ | The date the domain was created on the registry. | 
| dates_updated_at | timestamp | X | √ | The last update date. | 
| dates_hold_ends_at | timestamp | X | √ | The date from which the domain can’t be renewed anymore (the domain can be restored if the registry supports redemption period otherwise the domain might be destroyed at Gandi at that date). | 
| nameservers | json | X | √ | List of current nameservers. | 


