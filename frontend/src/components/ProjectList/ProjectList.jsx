import React from "react";
import "./ProjectList.scss";
import {UseDrag} from "react-dnd";

export default function ProjectList(props) {
    return ()
}

export default function Card({isDragging, text}) {
    const [{ opacity }, dragRef] = useDrag({
        item: { type: ItemTypes.CARD, text },
        collect: (monitor) => ({
        opacity: monitor.isDragging() ? 0.5 : 1
        })
    });

    return(
        <div ref={dragRef} style={{ opacity }}>
            {text}
        </div>
    )
}

function onReorder(e) {
    console.log("reordered");
    console.dir(e);
}

function itemClicked(e) {
    console.log("on click");
    console.dir(e);
}

function ListItem() {
    return (<div>Blah!</div>);
};