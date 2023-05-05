import styles from './App.module.css';
import {createSignal, For} from "solid-js";
import {DialogBox} from "./components/DialogBox";
import {MessageBox} from "./components/MessageBox";
import {HistoryBox} from "./components/HistoryBox";

const USERNAME = "JoJo";
const SERVER_URL = `http://${import.meta.env.VITE_SERVER_ADDRESS}`;
const [dialogs, setDialogs] = createSignal([]);
const [chats, setChats] = createSignal([]);
const [current_chat_id, setCurrentChatId] = createSignal(1);
const [current_algorithm, setCurrentAlgorithm] = createSignal("kmp");

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

function loadDialogs() {
    fetch(
        `${SERVER_URL}/message?chat_id=${current_chat_id()}&limit=100&page=1`
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
        setCurrentChatId(json["chat_id"]);
    });
}

function loadChats() {
    fetch(`${SERVER_URL}/chat/${USERNAME}`).then(response => {
        if (!response.ok) {
            return Promise.reject(response);
        }
        return response.json();
    }).then(async data => {
        if (data !== null) {
            let result = await data;
            setCurrentChatId(result[0]);
            setChats(prev => prev.concat(result));
            loadDialogs(current_chat_id());
        } else {
            createChat();
        }
    }).catch(error => {
        console.log(error)
    });
}

export function loadData() {
    logInOrSignUp();
    loadChats();
}

function sendMsg(msg) {
    console.log(`Send message: ${msg}`)
    fetch(`${SERVER_URL}/message`, {
        method: "POST",
        body: JSON.stringify({
            chat_id: current_chat_id(),
            question: msg,
            algorithm: current_algorithm(),
        })
    }).then(response => {
        if (!response.ok) {
            return Promise.reject(response);
        }
        setDialogs(prev => prev.concat([msg]));
        return response.json();
    }).then(async data => {
        let result = await data;
        if (data !== null) {
            setDialogs(prev => prev.concat([result["answer"]]));
        }
    })
}

function App() {
    loadData();

    return (
        <div className={styles.App}>
            <div className={styles.left}>
                <div style={{
                    height: "100%",
                    display: "flex",
                    "flex-direction": "column",
                }}>
                    <div style={{
                        position: "relative",
                        height: "100%",
                        width: "100%",
                        flex: "1 1 0",
                    }}>
                        <nav style={{
                            height: "100%",
                            width: "100%",
                            display: "flex",
                            padding: "0.5rem",
                            "flex-direction": "column",
                        }}>
                            <div role="button"
                                 style={{
                                     margin: "0.5rem",
                                     display: "flex",
                                     padding: "0.75rem",
                                     "align-items": "center",
                                     "flex-shrink": "0",
                                     color: "white",
                                     cursor: "pointer",
                                     "margin-bottom": "0.25rem",
                                     gap: "0.75rem",
                                     border: "1px solid #F1F3F4",
                                     "border-radius": "0.375rem",
                                     "font-size": "0.875rem",
                                     "line-height": "1,25rem",
                                 }}>
                                <span className="material-symbols-outlined" 
                                      style={{color: "white"}}>
                                    add_circle
                                </span>
                                New chat
                            </div>
                            <div style={{
                                flex: "1 1 0",
                                "overflow-y": "auto",
                            }}>
                                <div style={{
                                    display: "flex",
                                    "flex-direction": "column",
                                    gap: "0.5rem",
                                    "font-size": "0.875rem",
                                    "line-height": "1.25rem",
                                    "padding-bottom": "0.5rem",
                                }}>
                                    <div style={{position: "relative"}}>
                                        <ol>
                                            <For each={chats()}>{(chat, idx) =>
                                                <HistoryBox id={chat}
                                                            onClick={id => setCurrentChatId(id)}
                                                            selected={chat === current_chat_id()}>
                                                    {`Chat ${idx() + 1}`}
                                                </HistoryBox>
                                            }</For>
                                        </ol>
                                    </div>
                                </div>
                            </div>
                        </nav>
                    </div>
                </div>
            </div>
            <div className={styles.right}>
                <main className={styles.main}>
                    <div className={styles.dialogsContainer}>
                        <div style={{
                            position: "relative",
                            height: "100%",
                        }}>
                            <div style={{
                                height: "100%",
                                width: "100%",
                                "overflow-y": "auto",
                            }}>
                                <div style={{
                                    display: "flex",
                                    "flex-direction": "column",
                                    "align-items": "center",
                                }}>
                                    <For each={dialogs()}>{(dialog, idx) =>
                                        <DialogBox speaker={idx() % 2 === 0 ? `${USERNAME}` : "GPT-chan"}
                                                   background={idx() % 2 === 0 ? "" : "#313244"}>
                                            {dialog}
                                        </DialogBox>
                                    }</For>
                                    <div style={{
                                        width: "100%",
                                        height: "12rem",
                                        "flex-shrink": "0",
                                    }}></div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className={styles.formSection}>
                        <MessageBox onSend={sendMsg}/>
                        <div style={{
                            color: "white",
                            "padding-bottom": "1.75rem",
                            "padding-top": "0.75rem",
                            "padding-left": "1rem",
                            "padding-right": "1rem",
                            "font-size": "0.75rem",
                            "line-height": "1rem",
                            "text-align": "center",
                        }}>
                            Created by: Addin Munawwar, Moch. Sofyan Firdaus, and Ezra M. C. M. H.
                        </div>
                    </div>
                </main>
            </div>
        </div>
    );
}

export default App;
