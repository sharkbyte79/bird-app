# BirdUp
A full-stack web application for exploring bird observation data via the public [eBird API](https://ebird.org/home). Enables search for observations by **region code** (e.g. US-CT, US-NY). 

## Tech Stack
- **Backend:** Go ([Gin](https://github.com/gin-gonic/gin) framework)
- **Frontend:** TypeScript, React and Tailwind CSS
- **External Services**: Firebase (Authentication).

## Schema
- `User`: Models a single User's account registered on the Birdup site. 
    - `firebaseId` (pk): The unique identifier for this account assigned by Firebase's authentication services.
    - `email`: The email address provided at signup.
    - `createdAt`: The timestamp for the moment of this user's creation.

- `BirdFollow`: Models the relation between a single User and a bird species they've followed observations of.
    - `id` (pk): The unique identifier for this record.
    - `userId` (fk): The identifying Firebase ID assigned to the relevant User.
    - `speciesCode`: The unique identifier for the relevant bird species as assigned by eBird. 
    - `createdAt`: The timestamp for the moment this record was created.
