import {createSignal} from "solid-js";
import styles from "../App.module.css";

export function MessageBox(props) {
    const [msg, setMsg] = createSignal("");

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
            props.setDialogs(props.dialogs().concat([msg]));
            return response.json();
        }).then(async data => {
            let result = await data;
            if (data !== null) {
                props.setDialogs(props.dialogs().concat([result["answer"]]))
            }
        })
        setMsg("");
    }
    const questionKeyDown = event => {
        if (event.key === "Enter" && !event.shiftKey) {
            event.preventDefault();
            sendMsg(msg())
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
                        if (msg().length > 0) sendMsg(msg())
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