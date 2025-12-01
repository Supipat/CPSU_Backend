import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import axios from "axios";
import Headers from "../../component/header";
import Navbar from "../../component/navbar";
import Footer from "../../component/footer";
import Menu from "../../component/menu";
import { Link } from "react-router-dom";
import "../../css/admin/subject.css";

const Subject_detail = () => {
  const { id } = useParams(); // ดึง id จาก URL params
  const [subject, setSubject] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchSubjectDetail = async () => {
      try {
        const res = await axios.get(
          `http://localhost:8080/api/v1/admin/subject/${id}`
        );
        setSubject(res.data);
      } catch (err) {
        setError("ไม่พบข้อมูลรายวิชา หรือเกิดข้อผิดพลาดในการดึงข้อมูล");
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    fetchSubjectDetail();
  }, [id]);
  const handleDelete = async () => {
    const confirmDelete = window.confirm(
      `คุณต้องการลบรายวิชา ${subject.subject_id} ใช่หรือไม่?`
    );
    if (!confirmDelete) return;

    try {
      await axios.delete(`http://localhost:8080/api/v1/admin/subject/${id}`);
      alert("ลบรายวิชาสำเร็จ");
      navigate("/admin/subject");
    } catch (err) {
      alert("เกิดข้อผิดพลาดในการลบรายวิชา");
      console.error(err);
    }
  };

  if (loading) return <p>กำลังโหลดข้อมูล...</p>;
  if (error) return <p>{error}</p>;
  if (!subject) return <p>ไม่พบข้อมูลรายวิชา</p>;

  return (
    <>
      <Headers />
      <Navbar />
      <div className="container mt-4">
        <div className="row">
          <div className="col-sm-3">
            <Menu />
          </div>
          <div className="col-sm-9 text-start">
            <div className="d-flex justify-content-between align-items-center mb-3">
              {/* ซ้าย: subject name */}
              <p id="subject-name" className="m-0 fw-bold">
                {subject.subject_id} {subject.thai_subject}
              </p>

              <div>
                <Link
                  to={`/admin/editsubject/${subject?.id}`}
                  className="btn btn-warning me-2"
                  id="btn-edit"
                >
                  แก้ไข
                </Link>
                <button
                  className="btn btn-danger"
                  id="btn-delete"
                  onClick={handleDelete}
                >
                  ลบ
                </button>
              </div>
            </div>

            <p>
              <strong>หลักสูตร:</strong> {subject.thai_course}
            </p>
            <p>
              <strong>ภาคการศึกษา:</strong> {subject.semester}
            </p>
            <p>
              <strong>หน่วยกิต:</strong> {subject.credits}
            </p>

            <p>
              <strong>วิชาบังคับ:</strong> {subject.compulsory_subject || "-"}
            </p>

            <p>
              <strong>เงื่อนไข:</strong> {subject.condition || "-"}
            </p>
            <hr />
            <p id="subject-detail">คำอธิบายรายวิชา (ภาษาไทย)</p>
            <p>{subject.description_thai || "-"}</p>
            <p id="subject-detail">Course Description (English)</p>
            <p>{subject.description_eng || "-"}</p>
            <hr />
            <p id="subject-detail">CLOs</p>
            <pre
              style={{
                whiteSpace: "pre-wrap",
                backgroundColor: "#f8f9fa",
                padding: "15px",
                borderRadius: "5px",
              }}
            >
              {subject.clo || "-"}
            </pre>
          </div>
        </div>
      </div>
      <Footer />
    </>
  );
};

export default Subject_detail;
