import React from "react";
import "../css/menu.css";
import { Link, useLocation } from "react-router-dom";

const Menu = () => {
  const location = useLocation();
  const currentPath = location.pathname;

  const menuItems = [
    { path: "/admin/news", label: "ข่าวสาร", icon: "fas fa-newspaper" },
    { path: "/admin/personnel", label: "บุคลากร", icon: "fas fa-user-tie" },
    { path: "/admin/course", label: "หลักสูตร", icon: "fas fa-book-open" },
    { path: "/admin/subject", label: "รายวิชา", icon: "fas fa-book" },
    { path: "/admin/calendar", label: "ปฏิทินกิจกรรม", icon: "fas fa-calendar-alt" },
    { path: "/admin/userpermissions", label: "กำหนดสิทธิ์ผู้ใช้", icon: "fas fa-user-shield" },
  ];

  return (
    <div>
      <ul className="list-group">
        {menuItems.map((item) => (
          <Link
            key={item.path}
            to={item.path}
            className={`list-group-item ${currentPath === item.path ? "active" : ""}`}
            id="btn-menu"
          >
            <i className={`${item.icon} me-2`}></i>
            {item.label}
          </Link>
        ))}
      </ul>
    </div>
  );
};

export default Menu;
