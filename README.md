# Tugas Besar PAT 2023/2024

NIM | Nama
--- | --- 
13521114 | Farhan Nabil Suryono
13521148 | Johanes Lee
13521150 | I Putu Bakta Hari Sudewa

---

## Requirement Sistem

Terdapat kebutuhan akan perangkat lunak yang dapat digunakan untuk pemesanan tiket pada acara tertentu. Perangkat lunak ini akan digunakan sebagai penengah customer/client/app lain untuk melakukan booking. Karena perangkat lunak perlu mengakomodasi load yang besar dan concurrency yang tinggi, sistem yang dikembangkat perlu mengadaptasi arsitektur microservice yang setidaknya terdiri atas 3 buah service:
Client App atau service yang berinteraksi dengan end user,
Ticket App atau service yang menyimpan data kursi beserta ketersediaannya untuk setiap event, serta
Payment App atau service yang menyediakan payment gateway.


## Arsitektur Sistem
![architecture](/assets//architecture.png)

## Menjalankan Program

Ikuti readme pada masing-masing folder.

## Postman
### Booking
![Booking](/assets//postman1.jpg)
### Create Invoice
![Invoice](/assets//postman2.jpg)
### Payment
![Payment](/assets//postman3.jpg)

## Client
### Sign In
![Sign In](/assets//client1.jpg)
### Profile
![Profile](/assets//client2.jpg)
### History
![History](/assets//client3.jpg)

