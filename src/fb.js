import Firebase from "firebase/app";
import "firebase/firestore";

const opts = {
    apiKey: "AIzaSyA4BokV9JSpzry0bGB72Ap2QcxjQ-w3m7A",
    authDomain: "minstar-f0e35.firebaseapp.com",
    databaseURL: "https://minstar-f0e35.firebaseio.com",
    projectId: "minstar-f0e35",
    storageBucket: "minstar-f0e35.appspot.com",
    messagingSenderId: "121749245040",
    appId: "1:121749245040:web:efffab6f67aaf418a7b8b5",
    measurementId: "G-D53WMHPGQG",
};
var fbapp
if (Firebase.apps.length) {
    fbapp = Firebase.app()
} else {
    fbapp = Firebase.initializeApp(opts);
}

export const db = fbapp.firestore()
