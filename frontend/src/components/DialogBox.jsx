export function DialogBox(props) {
    return (
        <div style={{
            width: "100%",
            background: props.background,
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
                <div style={{
                    width: "4.5rem",
                    "max-width": "4.5rem",
                }}>
                    {props.speaker}
                </div>
                <div style={{
                    position: "relative",
                    width: "calc(100% - 100px)",
                }}>
                    <div style={{
                        "min-height": "20px",
                        "white-space": "pre-wrap",
                        "word-wrap": "break-word",
                    }}>
                        {props.children}
                    </div>
                </div>
            </div>
        </div>
    )
}