import {createSignal} from "solid-js";
import styles from "../App.module.css";

export function MessageBox(props) {
    const [msg, setMsg] = createSignal("");
    const questionKeyDown = event => {
        if (event.key === "Enter" && !event.shiftKey && msg().length > 0) {
            event.preventDefault();
            props.onSend(msg());
            setMsg("");
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
    const sendMsgEnter = ({target}) => {
        if (msg().length > 0) {
            target.style.background = "#11111180";
            target.style.cursor = "pointer";
        }
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
                    onClick={_ => {
                        if (msg().length > 0) {
                            props.onSend(msg());
                            setMsg("");
                        }
                    }}
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
