# Sky tech test

This project is a solution to the practical exercise provided by Popsa. It is a single service that reads in CSV files, each file is considered an album of photos metadata. We use the metadata to reverse geolocate where the photos were taken using here.coms API. Finally it generates appropriate titles based off this data. 

## Command Options

You can manage and run this solution using the following terminal commands from the repository root:

- `make run` - Run the service


## Implementation

### Code
The code follows idiomatic Go patterns consisting of an internal directory containing core application.


### Future changes / improvements 

#### Database
- No database has been used for this solution and I didn't think it was neccesary given the requirements. However in the future we would want to persist meta data, we would need more information such as a user the albums are associated with.
#### Code 

- AI to generate titles: We currently have a very trivial implementation for generating titles. AI would be a great solution for this, providing metadata such as location, weather and timeline it could generate bespoke titles 

- More in depth meta data: The amount of data is very limiting, we would want more in depth meta data in the future to provide more meaningful titles.

- User interaction: Feedback from the user, are the titles good, do they want something specific. Even a user inputted story to give more context, things such as inside jokes, specific events they want highlighting. 

