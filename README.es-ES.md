

# sms-nigeria-go — Envía SMS a cualquier número de teléfono de Nigeria con facilidad

[![Build Status](https://travis-ci.com/D-sense/sms-nigeria-go.svg?branch=master)](https://travis-ci.com/D-sense/sms-nigeria-go.svg?branch=master)

## Proveedores de SMS compatibles                   
| Plataforma          | ¿Integrado?        
| ----------------- | -------------       
| Bulk Sms Nigeria  | SÍ                  
| Sms Clone         | SÍ                  
| Net Bulk SMS      | NO                   
| Padisoft          | NO                   
| Sms mobile 24     | NO                   

## Instalación
go get github.com/d-sense/sms-nigeria-go


## Uso

> Envía SMS a través del servicio Bulk_Sms_Nigeria (primero, crea y activa una cuenta aquí para obtener las credenciales: https://www.bulksmsnigeria.com/)

```go
// import the package
github.com/d-sense/sms-nigeria-go

// declare the notification service
var bulkSmsNigeria sms_nigeria_go.SmsController

// declare, initialize, and send data
data := &sms_nigeria_go.BulkSmsNigeriaNotification{
	Sender:    os.Getenv("SMS_SENDER"),
	Recipient: recipient,
	Body:      textMessage,
	ApiToken:  os.Getenv("BULK_SMS_NIGERIA_API_TOKEN"),
}

result, err := repo.SendBulkSmsNigeria(data)
if err != nil {
	log.Fatal(err)
}
fmt.Println(result)
```


> Envía SMS a través del servicio Sms_Cone (primero, crea y activa una cuenta aquí para obtener las credenciales: http://smsclone.com)

```go
// import the package
github.com/d-sense/sms-nigeria-go

// declare the notification service
var smsClone sms_nigeria_go.SmsController

// declare, initialize, and send data
smsCloneData := &sms_nigeria_go.SmsCloneNotification{
	Username:  os.Getenv("USERNAME"),
	Password:  os.Getenv("PASSWORD"),
	Sender:    os.Getenv("SMS_SENDER"),
	Recipient: recipient,
	Message:   textMessage,
}

result, err := repo.SendSmsClone(smsCloneData, sms_nigeria_go.SmsCloneNormalRoute)
if err != nil {
	log.Fatal(err)
}
fmt.Println(result)

// to check credit balance
credential := sms_nigeria_go.SmsCloneCredential{
	Username: os.Getenv("USERNAME"),
	Password: os.Getenv("PASSWORD"),
}

balance, err := repo.CheckBalanceSmsClone(&credential)
if err != nil {
	log.Fatal(err)
}
fmt.Println(balance)
```

#### Ten en cuenta que para el servicio Sms_Clone, puedes usar diferentes métodos para enviar un SMS y asegurar que el mensaje llegue. Por mi experiencia, "dnd-fallback" parece ser el más fiable (aunque el costo varía según el método).
Método | Constante
:----: | --------
sendsms route | SmsCloneNormalRoute
dnd-route | SmsCloneDndRoute
dnd-fallback | SmsCloneNormalAndDndRoute
---


## Licencia

[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)
- **[Licencia MIT](http://opensource.org/licenses/mit-license.php)**
