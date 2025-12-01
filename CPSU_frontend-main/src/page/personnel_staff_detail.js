import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";
import Navbar from "../component/navbar";
import Headers from "../component/header";
import Footer from "../component/footer";
import "../css/people_staff_detail.css";

const TeacherDetail = () => {
  const { id } = useParams();
  const [person, setPerson] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [copied, setCopied] = useState(false);

  useEffect(() => {
    const fetchPerson = async () => {
      try {
        const res = await axios.get(
          `http://localhost:8080/api/v1/admin/personnel/${id}`
        );
        setPerson(res.data.personnel);
        setLoading(false);
      } catch (err) {
        console.error("Error fetching personnel detail:", err);
        setError(err);
        setLoading(false);
      }
    };

    fetchPerson();
  }, [id]);

  const handleCopyEmail = () => {
    if (person?.email) {
      navigator.clipboard.writeText(person.email);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    }
  };

  if (loading) {
    return (
      <>
        <Headers />
        <Navbar />
        <div className="container text-center my-5">
          <div className="spinner-border text-primary" role="status"></div>
          <p className="mt-3">กำลังโหลดข้อมูล...</p>
        </div>
        <Footer />
      </>
    );
  }

  if (error || !person) {
    return (
      <>
        <Headers />
        <Navbar />
        <div className="container my-5">
          <div className="alert alert-danger text-center">
            เกิดข้อผิดพลาดไม่สามารถโหลดข้อมูลได้
          </div>
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
        <div className="teacher-card mx-auto bg-white shadow rounded p-4">
          <div className="row">
            {/* Left: Image and Email */}
            <div className="col-md-4 text-center mb-4 mb-md-0">
              <img
                src={person.file_image}
                alt={person.thai_name}
                className="img-fluid img-profile"
                onError={(e) => (e.target.src = "/images/default-profile.png")}
              />

              <div className="d-flex justify-content-center align-items-center gap-2 mt-3 flex-wrap">
                <div>
                  <strong>อีเมล</strong>
                  <br />
                  <a href={`mailto:${person.email}`}>{person.email}</a>
                </div>
                <button
                  className="btn btn-outline-primary btn-sm d-flex align-items-center gap-1 copy-btn"
                  onClick={handleCopyEmail}
                  title="คัดลอกอีเมล"
                >
                  <i className="bi bi-clipboard"></i> คัดลอก
                </button>
              </div>
              {copied && (
                <small className="text-success d-block mt-1">คัดลอกแล้ว!</small>
              )}
            </div>

            {/* Right: Info */}
            <div className="col-md-8 text-start">
              <h5 className="text-primary">
                {person.department_position_name}
              </h5>
              <hr></hr>
              <h4 className="fw-bold mt-2">
                {person.thai_academic_position} {person.thai_name}
              </h4>
              <p className="text-muted mb-1">
                {person.eng_academic_position} {person.eng_name}
              </p>
              <hr></hr>
              <p className="mt-3">
                <strong>เว็บไซต์ </strong>{" "}
                <a href={person.website} target="_blank" rel="noreferrer">
                  {person.website ? person.website : "-"}
                </a>
              </p>
            </div>
          </div>
        </div>
      </div>

      <Footer />
    </>
  );
};

export default TeacherDetail;
