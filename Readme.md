# Popsa test

This project is a solution to the practical exercise provided by Popsa. It is a single service that processes CSV files, where each file represents an album of photo metadata. The service uses the metadata to reverse geolocate where the photos were taken via the Here.com API, then generates appropriate titles based on this data. 

## Command Options

You can manage and run this solution using the following terminal commands from the repository root:

- `make run` - Run the service
- `make test` - Run tests

## Implementation

### Code
The code follows idiomatic Go patterns consisting of an internal directory containing core application.


### Future changes / improvements 

#### Database
- No database was used in this solution as it was not necessary given the current requirements. However, in the future, we may want to persist metadata, which would require additional information, such as the user the albums are associated with.
#### Code 

- Batching API calls

- Logging and Error handling

- AI to generate titles: Currently, the title generation is fairly basic. In the future, AI could enhance this by using metadata such as location, weather, and timelines to generate more creative and personalized titles.

- Expanded Metadata: The current dataset is somewhat limited. Collecting more in-depth metadata in the future would allow for generating more meaningful and context-rich titles.

- User Interaction: Adding user feedback mechanisms would improve the title generation process. Users could provide input on whether the titles are accurate or relevant, or specify certain themes or events they want highlighted. This could include things like inside jokes, specific events, or personalized preferences.