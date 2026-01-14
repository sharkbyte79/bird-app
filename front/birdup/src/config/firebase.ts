interface FirebaseConfig {
    apiKey: string;
    authDomain: string;
    projectId: string;
    storageBucket: string;
    messagingSenderId: string;
    appId: string;
    measurementId?: string;
}

export const config: FirebaseConfig = {
    // NOTE api key is only used for identifying the Firebase project
    // and may be safely exposed (https://firebase.google.com/docs/projects/api-keys)
    apiKey: "AIzaSyCZ81hSYOXCTtwv0_mbJ1-xUp042usl19Y",
    authDomain: "birdup-ec859.firebaseapp.com",
    projectId: "birdup-ec859",
    storageBucket: "birdup-ec859.firebasestorage.app",
    messagingSenderId: "345444436926",
    appId: "1:345444436926:web:bd1f204e6fdb76b656c7ef",
    measurementId: "G-HDRRJS70B9"
}

