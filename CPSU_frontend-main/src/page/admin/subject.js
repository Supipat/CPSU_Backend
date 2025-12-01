// ปรับปรุงแล้ว: เพิ่ม className, ปรับ UX และคำอธิบายให้โค้ดอ่านง่ายขึ้น
import React, { useEffect, useState } from "react";
import "../../css/admin/subject.css";
import { Link, useNavigate } from "react-router-dom";
import axios from "axios";
import Headers from "../../component/header";
import Navbar from "../../component/navbar";
import Footer from "../../component/footer";
import Menu from "../../component/menu";

function Subject() {
  const [subject, setSubject] = useState([]);
  const [searchTerm, setSearchTerm] = useState("");
  const [showDropdown, setShowDropdown] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchSubject = async () => {
      try {
        const res = await axios.get(
          "http://localhost:8080/api/v1/admin/subject"
        );
        const filtered = res.data.filter((item) => {
          const sid = item.subject_id?.trim();
          const name = item.thai_subject?.trim();
          return (
            sid !== "------" &&
            name !== "วิชาเลือกเสรี" &&
            name !== "วิชาเลือกในหมวดวิชาเฉพาะ" &&
            sid !== "SUxxx" &&
            sid !== "SUXXX"
          );
        });
        setSubject(filtered);
      } catch (err) {
        console.error("Error fetching subject", err);
      }
    };

    fetchSubject();
  }, []);

  const filteredSubjects = subject.filter((item) => {
    const term = searchTerm.toLowerCase();
    return (
      item.subject_id.toLowerCase().includes(term) ||
      item.thai_subject.toLowerCase().includes(term)
    );
  });

  const groupedByCourse = subject.reduce((acc, item) => {
    const courseId = item.course_id;
    if (!acc[courseId]) {
      acc[courseId] = [];
    }
    acc[courseId].push(item);
    return acc;
  }, {});

  const handleSelect = (id) => {
    setShowDropdown(false);
    setSearchTerm("");
    navigate(`/admin/subject/${id}`);
  };

  return (
    <>
      <Headers />
      <Navbar />
      <div className="container text-center mt-4">
        <div className="row">
          <div className="col-sm-3">
            <Menu />
          </div>

          <div className="col-sm-9">
            <div className="d-flex justify-content-between align-items-center mb-4">
              <h3 className="subject-title">รายวิชาทั้งหมด</h3>
              <Link to="/admin/addsubject" className="btn-addsubject">
                 + เพิ่มรายวิชาใหม่
              </Link>
            </div>

            <div className="searchbox-subject">
              <input
                type="text"
                placeholder="ค้นหารหัสวิชา หรือ ชื่อวิชา..."
                value={searchTerm}
                onChange={(e) => {
                  setSearchTerm(e.target.value);
                  setShowDropdown(true);
                }}
                onBlur={() => setTimeout(() => setShowDropdown(false), 150)}
                onFocus={() => {
                  if (searchTerm.length > 0) setShowDropdown(true);
                }}
                className="form-control"
              />

              {/* Dropdown search suggestions */}
              {showDropdown && searchTerm && (
                <ul className="dropdown-menu show custom-dropdown">
                  {filteredSubjects.length > 0 ? (
                    filteredSubjects.map((item) => (
                      <li key={item.id}>
                        <button
                          type="button"
                          className="dropdown-item"
                          onMouseDown={() => handleSelect(item.id)}
                        >
                          {item.subject_id} - {item.thai_subject}
                        </button>
                      </li>
                    ))
                  ) : (
                    <li className="dropdown-item disabled">ไม่พบผลลัพธ์</li>
                  )}
                </ul>
              )}
            </div>

            {/* แสดงตารางรายวิชา */}
            {Object.keys(groupedByCourse).map((courseId) => (
              <div key={courseId} className="mb-4">
                <h5 className="text-start course-heading">หลักสูตร: {courseId}</h5>
                <div className="table-responsive">
                  <table className="table table-bordered text-start">
                    <thead className="table-light">
                      <tr>
                        <th style={{ width: "15%" }}>รหัสวิชา</th>
                        <th>ชื่อวิชา</th>
                        <th style={{ width: "15%" }}>หน่วยกิต</th>
                      </tr>
                    </thead>
                    <tbody>
                      {groupedByCourse[courseId].map((subj) => (
                        <tr key={subj.id}>
                          <td>
                            <Link to={`/admin/subject/${subj.id}`}>
                              {subj.subject_id}
                            </Link>
                          </td>
                          <td>{subj.thai_subject}</td>
                          <td>{subj.credits}</td>
                        </tr>
                      ))}
                    </tbody>
                  </table>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
      <Footer />
    </>
  );
}

export default Subject;
