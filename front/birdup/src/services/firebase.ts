import { FirebaseApp, initializeApp } from "firebase/app";
import { getAuth, GoogleAuthProvider } from "firebase/auth";
import type { Auth } from "firebase/auth";
import { config } from "../config/firebase";

// Initialize Firebase
const app: FirebaseApp = initializeApp(config);

// Initialize Firebase Authentication and get a reference to the service
const auth: Auth = getAuth(app);

// Initialize authentication via Google and get a reference to the provider
const googleProvider: GoogleAuthProvider = new GoogleAuthProvider();

export { app, auth, googleProvider };

