import styles from './App.module.css';
import {createSignal, For} from "solid-js";
import {DialogBox} from "./components/DialogBox";
import {MessageBox} from "./components/MessageBox";

function App() {
    const [dialogList, setDialogs] = createSignal(dialogs);


    const sendMsg = msg => {
        console.log(`Send message: ${msg}`)
        fetch(`${SERVER_URL}/message`, {
            method: "POST",
            body: JSON.stringify({
                chat_id: current_chat_id,
                question: msg,
                algorithm: current_algorithm,
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
                setDialogs(prev => prev.concat([result["answer"]]))
            }
        })
    }

    return (
        <div className={styles.App}>
            <div className={styles.left}></div>
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
                                    <For each={dialogList()}>{(dialog, idx) =>
                                        <DialogBox speaker={idx() % 2 === 0 ? `${USERNAME}` : "GPT-chan"} background={idx() % 2 === 0 ? "" : "#313244"}>
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
