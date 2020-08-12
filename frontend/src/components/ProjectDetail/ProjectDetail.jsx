import React from "react";
import "./ProjectDetail.scss";
import useParams from "react";

export default function ProjectDetail() {
    let { id } = useParams();
    return (<div>Project ID: {id}</div>);
}