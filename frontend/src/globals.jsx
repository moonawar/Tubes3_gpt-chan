import {createSignal} from "solid-js";

// global variables
export const USERNAME = "JoJo";
export const SERVER_URL = `http://${import.meta.env.VITE_SERVER_ADDRESS}`;
export const [dialogs, setDialogs] = createSignal([]);

export let current_algorithm = "kmp";
export let current_chat_id = 1;

function logInOrSignUp() {
    fetch(`${SERVER_URL}/user`, {
        method: "POST",
        body: JSON.stringify({username: USERNAME}),
    }).then(response => {
        if (!response.ok) {
            return Promise.reject(response);
        }
        return response.json();
    }).then(_ => {
        console.log(`Sign Up: ${USERNAME}`)
    }).catch(_ => {
        console.log(`Sign In: ${USERNAME}`)
    });
}

function loadDialogs(chat_id) {
    fetch(
        `${SERVER_URL}/message?chat_id=${current_chat_id}&limit=100&page=1`
    ).then(response => {
        if (!response.ok) {
            return Promise.reject(response);
        }
        return response.json();
    }).then(async data => {
        let arr = await data;
        if (arr !== null) {
            for (let obj of Object.values(arr)) {
                setDialogs(prev => prev.concat(obj["question"], obj["answer"]));
            }
        }
    }).catch(error => {
        console.log(error);
    });
}

function createChat() {
    fetch(`${SERVER_URL}/chat`, {
        method: "POST",
        body: JSON.stringify({
            username: USERNAME
        }),
    }).then(response => {
        if (!response.ok) {
            return Promise.reject(response);
        }
        return response.json();
    }).then(async data => {
        let json = await data;
        current_chat_id = json["chat_id"];
    });
}

function loadChat() {
    fetch(`${SERVER_URL}/chat/${USERNAME}`).then(response => {
        if (!response.ok) {
            return Promise.reject(response);
        }
        return response.json();
    }).then(async data => {
        if (data !== null) {
            let result = await data;
            current_chat_id = result[0];
            loadDialogs(current_chat_id);
        } else {
            createChat();
        }
    }).catch(error => {
        console.log(error)
    });
}

export function loadData() {
    logInOrSignUp();
    loadChat();
}