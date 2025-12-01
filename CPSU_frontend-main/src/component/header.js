import React from "react";
import "../css/header.css";
import Theme from "./theme";
import ProfileMenu from "../component/ProfileMenu";
import { Link } from "react-router-dom";

const Headers = () => {
  return (
    <>
      <div>
        <div className="header-component">
          <ProfileMenu />
          <div class="d-grid gap-2 d-md-flex justify-content-md-end">
            <Theme />
            {/* <Link to={`/login`}>
              <button className="btn btn-primary" type="button">
                เข้าสู่ระบบ
              </button>
            </Link> */}
          </div>
        </div>
      </div>
    </>
  );
};

export default Headers;
