# JSCroot Mobile

Support Capacitor

### **🔹 Cara Menggunakan Capacitor Tanpa Android Studio**
Untuk menghindari penggunaan Android Studio, kita bisa:
1. **Menggunakan SDK Android + Gradle secara manual**
2. **Menggunakan alat CLI untuk membangun dan menjalankan aplikasi**
3. **Menggunakan perangkat fisik atau emulator tanpa Android Studio**

---

## **1️⃣ Install Android SDK dan Gradle**
Karena Android Studio biasanya menyediakan **Android SDK dan Gradle secara otomatis**, kita perlu menginstalnya secara manual.

### **🔹 Install Android SDK Tanpa Android Studio**
1. **Unduh Android Command Line Tools:**  
   👉 [https://developer.android.com/studio#command-tools](https://developer.android.com/studio#command-tools)  
   ![image](https://github.com/user-attachments/assets/b1f98488-d97f-439a-99e4-3b614031aeaa)  
2. **Ekstrak dan tambahkan ke PATH:**  
   ```sh
   export ANDROID_HOME=$HOME/Android/Sdk
   export PATH=$ANDROID_HOME/cmdline-tools/latest/bin:$PATH
   export PATH=$ANDROID_HOME/platform-tools:$PATH
   ```
3. **Install platform SDK yang dibutuhkan:**  
   ```sh
   sdkmanager --install "platforms;android-33" "build-tools;33.0.2"
   sdkmanager --licenses
   ```

---

### **2️⃣ Buat Proyek Capacitor Tanpa npm/npx**
Jika Anda ingin **menghindari npm/npx**, Anda bisa **mengunduh Capacitor secara manual** dari CDN dan menambahkan file web ke proyek Android.

#### **📌 Struktur Folder**
```
/myApp
 ├── www/         # Folder untuk HTML, CSS, JS
 │   ├── index.html
 │   ├── script.js
 │   └── style.css
 ├── android/     # Folder proyek Android
 │   ├── app/
 │   ├── build.gradle
 │   └── settings.gradle
 └── capacitor.config.json
```

#### **📌 Buat `capacitor.config.json`**
```json
{
  "appId": "com.example.myapp",
  "appName": "My Capacitor App",
  "webDir": "www",
  "bundledWebRuntime": false
}
```

---

### **3️⃣ Build dan Jalankan Aplikasi Tanpa Android Studio**
Setelah konfigurasi selesai, kita bisa langsung membangun dan menjalankan aplikasi.

#### **🔹 Build APK menggunakan Gradle**
1. **Masuk ke folder proyek Android:**  
   ```sh
   cd android
   ```
2. **Jalankan Gradle untuk membangun aplikasi:**  
   ```sh
   ./gradlew assembleDebug
   ```
3. **Hasil APK akan ada di:**  
   ```
   android/app/build/outputs/apk/debug/app-debug.apk
   ```

---

### **4️⃣ Jalankan APK di Perangkat atau Emulator**
#### **🔹 Opsi 1: Menggunakan ADB**
Jika Anda memiliki perangkat Android atau emulator yang sudah berjalan, instal APK dengan:
```sh
adb install android/app/build/outputs/apk/debug/app-debug.apk
```

#### **🔹 Opsi 2: Menggunakan Emulator Ringan (Tanpa Android Studio)**
Alternatifnya, Anda bisa menggunakan emulator seperti **Genymotion** atau **Scrcpy** untuk menjalankan APK.

---

## **🔹 Kesimpulan**
✔ **Tanpa Android Studio**, Anda tetap bisa membangun aplikasi dengan Capacitor  
✔ **Gunakan Android SDK + Gradle secara manual** untuk compile APK  
✔ **Gunakan ADB atau emulator ringan** untuk menjalankan aplikasi  
✔ **Lebih ringan & fleksibel**, tapi butuh setup tambahan  

👉 **Metode ini cocok untuk yang ingin menghindari IDE berat seperti Android Studio!** 🚀
