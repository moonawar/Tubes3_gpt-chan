import styles from './App.module.css';
import {createSignal, For} from "solid-js";


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

function DialogBox(props) {
    return (
        <div style={{
            width: "100%",
        }}>
            <div style={{
                "font-size": "1rem",
                "line-height": "1.5rem",
                display: "flex",
                margin: "auto",
                gap: "1.5rem",
                "max-width": "48rem",
                "padding": "1.5rem 0",
            }}>
                <div>{props.speaker}</div>
                <div style={{
                    position: "relative",
                    width: "calc(100% - 100px)",
                }}>
                    <div style={{
                        "min-height": "20px",
                        "white-space": "pre-wrap",
                        "word-wrap": "break-word",
                    }}>
                        {props.text}
                    </div>
                </div>
            </div>
        </div>
    )
}

function App() {
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
                                    <For each={[...Array(10).keys()]}>{(x, _) =>
                                        <DialogBox speaker="You" text="Lorem ipsum dolor sit amet, consectetur adipisicing elit. Cupiditate eveniet nihil nostrum quae quis ratione, tempora vero? A aliquam, consequatur error fugit maxime porro veniam! Ad alias assumenda at eos, eveniet ex incidunt ipsum laborum maiores modi nemo nobis provident quis, quos repellendus totam ullam? Aliquam aperiam ducimus facere in minima necessitatibus nemo neque non omnis quas, quibusdam, reiciendis sed sint tempora veritatis. Aliquam animi assumenda cumque cupiditate excepturi facilis fugit hic in libero minus nam non nulla officia officiis possimus quae quas quo quos recusandae reiciendis sint soluta temporibus, unde ut vel vitae voluptatibus. Aliquid consequuntur dolores laboriosam nihil."/>
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
                        <MessageBox/>
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
