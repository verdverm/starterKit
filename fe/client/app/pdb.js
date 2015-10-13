import PouchDB from 'pouchdb';
// import { persistentStore } from 'redux-pouchdb';

let PDB = new PouchDB('app_local');
let remoteCouch = false;

export default PDB;