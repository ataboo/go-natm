import React from "react";
import "./userpanel.scss";
import { User } from "../../../../models/user";
import { IMainActions } from "../../imainactions";

type UserPanelProps = {
    mainActions: IMainActions
}

export const UserPanel = ({mainActions}: UserPanelProps) => { 
    return (<ul className="navbar-nav">
                <li className="align-items-center">
                    <div>{mainActions.currentUser?.name ?? "No User"}</div>
                </li>
                <li className="align-items-center ml-3">
                    <button className="btn btn-secondary" onClick={mainActions.logout}>Logout</button>
                </li>
            </ul>);
};
