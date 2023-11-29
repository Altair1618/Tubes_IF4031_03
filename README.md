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

## Seeding Data (Khusus Ticket Service)
Akses database menggunakan adminer pada localhost:8888  
System : PostgreSQL  
Server : ticket_service_db  
Username : [sesuai env pada ticket service]    
password : [sesuai env pada ticket service]  
database : [sesuai env pada ticket service]  

Dapat menggunakan aplikasi lain selain adminer.  

Jalankan SQL pada folder Ticket_Service/schema/dummy.sql   

## Postman
### Booking
![Booking](/assets//postman1.jpg)
### Create Invoice
![Invoice](/assets//postman2.jpg)
### Payment
![Payment](/assets//postman3.jpg)

## Client
### Sign In
![Sign In](/assets//client1.png)
### Profile
![Profile](/assets//client2.png)
### History
![History](/assets//client3.png)

