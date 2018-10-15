# tw-challenge
> conference track management

```bash
# Run the application for the selected file path
npm run app talks.txt
npm run app ${your-file-path}

# Install dependencies for the test
npm install

# Run the tests for the application elements
npm run test
```

**Solution:**
- I built a simple solution without many layers of abstraction because I wanted to keep the code straightforward
- The core of the application is contained within the findTalk*() functions in the Track object
- The function returns the index in the Talks array where the best combination is found
- To do so, I take into account the remaining time to complete the session
