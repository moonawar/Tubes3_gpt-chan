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
        target.style.height = Math.max(target.scrollHeight, 20) + 'px';
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
        <div className={styles.msgWrapper}>
                <textarea
                    className={styles.msg}
                    id="question"
                    name="question"
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
    );
}

function App() {
    return (
        <div class={styles.App}>
            <MessageBox />
        </div>
    );
}

export default App;
