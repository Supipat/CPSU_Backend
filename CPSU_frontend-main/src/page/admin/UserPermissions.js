import React, { useState } from "react";
import Headers from "../../component/header";
import Menu from "../../component/menu";
import Footer from "../../component/footer";
import Navbar from "../../component/navbar";
import "../../css/admin/user.css";

import { Link } from "react-router-dom";

const usersMock = [
  {
    id: 1,
    name: "กานต์พิชชา ธรรมรักษ์",
    username: "Karnpitcha",
    role: "teacher",
  },
  {
    id: 2,
    name: "ธีรภัทร วงศ์กิตติชัย",
    username: "Teerapat",
    role: "teacher",
  },
  {
    id: 3,
    name: "ณัฐวุฒิ ศรีจันทร์",
    username: "Nattawut",
    role: "teacher",
  },
  {
    id: 4,
    name: "ปิยะนุช วัฒนสุนทร",
    username: "Piyanuch",
    role: "teacher",
  },
  {
    id: 5,
    name: "วิชญ์พล พรหมพิทักษ์",
    username: "Witphol",
    role: "student",
  },
];

const roleOptions = [
  // { value: "admin", label: "Admin" },
  { value: "teacher", label: "อาจารย์" },
  { value: "student", label: "นักศึกษา" },
];

function UserPermissions() {
  const [users, setUsers] = useState(usersMock);

  const updateUserRole = (userId, newRole) => {
    setUsers((prev) =>
      prev.map((user) =>
        user.id === userId ? { ...user, role: newRole } : user
      )
    );
  };

  const handleSave = () => {
    console.log("บันทึกสิทธิ์ผู้ใช้:", users);
    alert("บันทึกสิทธิ์ผู้ใช้เรียบร้อย!");
  };

  return (
    <div>
      <Headers />
      <Navbar />
      <div className="container text-center">
        <div className="row">
          <div className="col-sm-3">
            <Menu />
          </div>
          <div className="col-sm-9 ">
            <div className="row">
              <div className="d-flex justify-content-between align-items-center mb-4">
                <h3 className="user-title">กำหนดสิทธิ์ผู้ใช้</h3>
                <Link to="/admin/user">+ เพิ่มผู้ใช้</Link>
              </div>

              <table className="table table-bordered">
                <thead className="table-light text-center">
                  <tr>
                    <th>รายชื่อบุคลากร</th>
                    <th>Username</th>
                    <th>Password</th>
                    <th>สิทธิ์การเข้าถึง</th>
                  </tr>
                </thead>
                <tbody>
                  {users.map((user) => (
                    <tr key={user.id} className="align-middle">
                      <td>{user.name}</td>
                      <td>{user.username}</td>
                      <td>********</td>
                      <td>
                        <select
                          className="form-select"
                          value={user.role}
                          onChange={(e) =>
                            updateUserRole(user.id, e.target.value)
                          }
                        >
                          {roleOptions.map((role) => (
                            <option key={role.value} value={role.value}>
                              {role.label}
                            </option>
                          ))}
                        </select>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>

              <div className="text-center mt-4">
                <button className="btn btn-primary" onClick={handleSave}>
                  บันทึกสิทธิ์
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
}

export default UserPermissions;
