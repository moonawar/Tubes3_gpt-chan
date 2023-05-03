import styles from './App.module.css';
import {createSignal} from "solid-js";


function MessageBox() {
    const [msg, setMsg] = createSignal("");
    const sendMsgEnter = ({target}) => {
        if (msg().length > 0) {
            target.style.background = "#11111180";
            target.style.cursor = "pointer";
        }
    };
    const questionKeyDown = event => {
        if (event.key === "Enter" && !event.shiftKey) {
            event.preventDefault();
        }
    };
    const onMsgInput = ({target}) => {
        setMsg(target.value);
        target.style.height = 0;
        target.style.height = Math.max(target.scrollHeight, 24) + 'px';
        if (target.value.length > 0) {
            document.getElementById("send-msg-wrapper").style.color = "#DDD";
        } else {
            document.getElementById("send-msg-wrapper").style.color = "#AAA";
        }
    };
    const sendMsgLeave = ({target}) => {
        target.style.background = "transparent";
        target.style.cursor = "default";
    };

    return (
        <form className={styles.msgWrapper}>
            <div className={styles.textAreaWrapper}>
                <textarea
                    autoFocus={true}
                    className={styles.msg}
                    id="question"
                    name="question"
                    tabIndex={0}
                    value={msg()}
                    placeholder="Send a message."
                    onInput={onMsgInput}
                    rows={1}
                    onKeyDown={questionKeyDown}/>
                <div
                    id="send-msg-wrapper"
                    className={styles.sendMsg}
                    onMouseEnter={sendMsgEnter}
                    onMouseLeave={sendMsgLeave}
                    role="button">
                    <span id="send-msg"
                          className="material-symbols-outlined">
                        send
                    </span>
                </div>
            </div>
        </form>
    );
}

function App() {
    return (
        <div className={styles.App}>
            <div className={styles.left}></div>
            <div className={styles.right}>
                <main className={styles.main}>
                    <div className={styles.dialogsContainer}></div>
                    <div className={styles.formSection}>
                        <MessageBox/>
                        <div style={{
                            color: "white",
                            "padding-bottom": "1.75rem",
                            "padding-top": "0.75rem",
                            "padding-left": "1rem",
                            "padding-right": "1rem",
                            "font-size": "0.75rem",
                            "line-height": "1rem",
                            "text-align": "center"
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
