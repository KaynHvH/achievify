## Achievify - Goal Achiever Application

### Description
Achievify is a web application that helps users outline actionable steps to achieve their goals. It utilizes OpenAI's GPT-3.5 Turbo to generate step-by-step instructions based on the user's provided goal. The generated responses are stored and can be accessed through a dedicated response page.

### Features
- **Goal Generation**: Users can input their desired goals through a simple user interface.
- **AI Assistance**: The application employs GPT-3.5 Turbo to generate detailed steps to accomplish the user's goals.
- **Response Storage**: The generated responses are stored in a database, allowing users to revisit and track their goals.
- **Interactive Response Page**: Users can view the generated steps in an interactive checklist format.

### Technologies Used
- **Backend**: Go (Golang)
- **Web Framework**: Gorilla Mux
- **Database**: SQLite
- **AI Integration**: OpenAI GPT-3.5 Turbo
- **Frontend**: HTML, CSS, JavaScript

### Todo
- JWT Authentication
- User registration and login
- User's achieves
- Notes for each achievements
- Possibility of adding another goal on the response page
- Better UI

### Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/KaynHvH/achievify.git
    cd achievify
    ```
2. Install dependencies:
    ```bash
    go get -u github.com/gorilla/mux
    go get -u github.com/go-resty/resty/v2
    go get -u github.com/joho/godotenv
    go get -u github.com/google/uuid
    go get -u github.com/mattn/go-sqlite3
    go get -u golang.org/x/net
    ```
3. Create a `.env` file in the root directory with the following content:
    ```env
    TOKEN=YOUR_OPENAI_API_KEY
    ```
   Replace `YOUR_OPENAI_API_KEY` with your actual OpenAI API key.


4. Run the application in terminal by command:
    ```bash
    make
    ```
   Now you can open index.html file located in "static" folder.

### Usage
1. Access the application in a web browser.
2. Enter your desired goal in the input field and click "I want to achieve that!".
3. The AI will generate step-by-step instructions to achieve the goal.
4. The generated response ID will be displayed, and you can click on the provided link to view the detailed steps.
5. The response page will show the steps as checkboxes, allowing you to interactively track your progress.

### Preview
![Main page](/assets/Main%20page.png)
![Terminal](/assets/Terminal.png)
![Response page](/assets/Response%20page.png)

### License
This project is licensed under the [MIT License](LICENSE).
