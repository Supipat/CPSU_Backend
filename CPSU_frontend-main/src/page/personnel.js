import React, { useEffect, useState } from "react";
import axios from "axios";
import Navbar from "../component/navbar";
import Headers from "../component/header";
import Footer from "../component/footer";
import "../css/people.css";
import { Link } from "react-router-dom";

const Personnel = () => {
  const [personnelList, setPersonnelList] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [activeTab, setActiveTab] = useState("สายวิชาการ");

  // ประเภทบุคลากร
  const personnelTypes = ["สายวิชาการ", "สายสนับสนุนวิชาการ"];

  useEffect(() => {
    const fetchPersonnel = async () => {
      try {
        const response = await axios.get(
          "http://localhost:8080/api/v1/admin/personnel"
        );
        setPersonnelList(response.data);
        setLoading(false);
      } catch (err) {
        console.error("Axios error:", err);
        setError(err);
        setLoading(false);
      }
    };

    fetchPersonnel();
  }, []);

  const filteredPersonnel = personnelList.filter(
    (person) => person.type_personnel === activeTab
  );

  if (loading || error) {
    return (
      <>
        <Headers />
        <Navbar />
        <div className="container text-center my-5">
          {loading ? (
            <p>กำลังโหลดข้อมูล...</p>
          ) : (
            <p>เกิดข้อผิดพลาด: {error.message}</p>
          )}
        </div>
        <Footer />
      </>
    );
  }

  return (
    <>
      <Headers />
      <Navbar />
      <div className="container my-5">
        <h3 className="text-center mb-4">บุคลากร</h3>

        {/* Tabs */}
        <ul className="nav nav-tabs justify-content-center mb-4">
          {personnelTypes.map((type) => (
            <li className="nav-item" key={type}>
              <button
                className={`nav-link ${activeTab === type ? "active" : ""}`}
                onClick={() => setActiveTab(type)}
              >
                {type}
              </button>
            </li>
          ))}
        </ul>

        {/* Personnel Cards */}
        <div className="row">
          {filteredPersonnel.length > 0 ? (
            filteredPersonnel.map((person) => {
              // ตรวจสอบว่าเป็นอาจารย์หรือเจ้าหน้าที่
              const detailPath =
                person.type_personnel === "สายวิชาการ"
                  ? `/teacher/${person.personnel_id}`
                  : `/staff/${person.personnel_id}`;

              return (
                <div
                  className="col-lg-3 col-md-4 col-sm-6 mb-4"
                  key={person.personnel_id}
                >
                  <Link
                    to={detailPath}
                    className="text-decoration-none text-dark"
                  >
                    <div className="card person-card shadow-sm h-100">
                      <img
                        src={person.file_image}
                        loading="lazy"
                        className="card-img-top"
                        alt={person.thai_name}
                        onError={(e) =>
                          (e.target.src = "/images/default-profile.png")
                        }
                      />
                      <div className="card-body text-center">
                        <h5 className="card-title">
                          {person.thai_academic_position} {person.thai_name}
                        </h5>
                        <p className="card-text text-center">
                          {person.department_position_name}
                        </p>
                      </div>
                    </div>
                  </Link>
                </div>
              );
            })
          ) : (
            <div className="col-12 text-center">
              <p>ไม่พบบุคลากรในประเภทนี้</p>
            </div>
          )}
        </div>
      </div>
      <Footer />
    </>
  );
};

export default Personnel;
