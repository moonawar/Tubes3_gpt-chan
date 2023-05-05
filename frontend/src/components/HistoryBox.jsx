export function HistoryBox(props) {
    let bg = props.selected ? "#CCCCCC20" : "transparent";
    return (
        <div role="button"
             style={{
                 margin: "0.5rem",
                 position: "relative",
                 display: "flex",
                 padding: "0.75rem",
                 "align-items": "center",
                 gap: "0.75rem",
                 "border-radius": "0.375rem",
                 cursor: "pointer",
                 "word-break": "break-all",
                 background: bg,
             }}
             onMouseEnter={event => {
                 if (!props.selected) {
                     event.currentTarget.style.background = "#CCCCCC10";
                 }
             }}
             onMouseLeave={event => {
                 event.currentTarget.style.background = bg;
             }}>
            <span className="material-symbols-outlined"
                  style={{color: "white"}}>
                chat
            </span>
            <div style={{
                position: "relative",
                flex: "1 1 0",
                overflow: "hidden",
                "max-height": "1.25rem",
                color: "white",
            }}>
                {props.children}
            </div>
        </div>
    )
}