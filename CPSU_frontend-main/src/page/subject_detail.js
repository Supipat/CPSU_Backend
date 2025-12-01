import React, { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";
import axios from "axios";
import Navbar from "../component/navbar";
import Headers from "../component/header";
import Footer from "../component/footer";

const Subject_detail = () => {
  const { subject_id } = useParams();
  const [subject, setSubject] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchSubject = async () => {
      setLoading(true);
      setError(null);

      try {
        const res = await axios.get(
          `http://localhost:8080/api/v1/admin/subject`
        );
        const foundSubject = res.data.find(
          (subj) =>
            String(subj.id) === subject_id || subj.subject_id === subject_id
        );

        if (!foundSubject) {
          setError("ไม่พบข้อมูลรายวิชาที่ต้องการ");
          setSubject(null);
        } else {
          setSubject(foundSubject);
        }
      } catch (err) {
        setError("เกิดข้อผิดพลาดในการดึงข้อมูล");
        setSubject(null);
      } finally {
        setLoading(false);
      }
    };

    fetchSubject();
  }, [subject_id]);

  if (loading)
    return (
      <>
        <Headers />
        <Navbar />
        <div className="container text-center my-5">
          <div className="spinner-border text-primary" role="status"></div>
          <p className="mt-3 fw-semibold fs-5">กำลังโหลดข้อมูลรายวิชา...</p>
        </div>
        <Footer />
      </>
    );

  if (error)
    return (
      <>
        <Headers />
        <Navbar />
        <div className="container text-center my-5">
          <h3 className="text-danger mb-3">{error}</h3>
          <Link
            to={`/course/${subject.course_id}?tab=subjects`}
            className="btn btn-secondary mb-4 back-button"
            id="btn-back-to-course"
          >
            &larr; กลับ
          </Link>
        </div>
        <Footer />
      </>
    );

  if (!subject) return null;

  return (
    <>
      <Headers />
      <Navbar />
      <div className="container my-5">
        <div className="card shadow-sm p-4 text-start">
          <Link
            to={`/course/${subject.course_id}?tab=subjects`}
            className="btn btn-secondary mb-4 back-button"
            id="btn-back-to-course"
          >
            &larr; กลับ
          </Link>

          {/* ชื่อภาษาไทย */}
          <h3 className="text-primary mb-1">
            ({subject.subject_id}) {subject.thai_subject}
          </h3>

          {/* ชื่อภาษาอังกฤษ */}
          <h5 className="text-secondary fst-italic mb-4">
            {subject.eng_subject}
          </h5>

          {/* หน่วยกิต */}
          <div className="mb-3">
            <strong>หน่วยกิต </strong> {subject.credits}
          </div>

          {/* เงื่อนไขรายวิชา */}

          <div className="mb-3">
            <strong>เงื่อนไขรายวิชา </strong>
            <p>{subject.condition || "-"}</p>
          </div>

          {/* วิชาบังคับก่อน */}
          <div className="mb-3">
            <strong>วิชาบังคับก่อน </strong>
            <p>{subject.compulsory_subject || "-"}</p>
          </div>

          <hr />

          {/* คำอธิบายรายวิชา */}
          <div className="mb-4">
            <h5>คำอธิบายรายวิชา</h5>
            <p>
              <strong>ภาษาไทย:</strong> {subject.description_thai}
            </p>
            <p>
              <strong>ภาษาอังกฤษ:</strong> {subject.description_eng}
            </p>
          </div>

          {/* ผลลัพธ์การเรียนรู้ (CLO) */}
          {/* ผลลัพธ์การเรียนรู้ (CLO) */}
          {subject.clo && (
            <div>
              <h5>ผลลัพธ์การเรียนรู้ (CLO)</h5>
              {subject.clo.split("\n").map(
                (line, index) =>
                  line.trim() !== "" && (
                    <p key={index} style={{ whiteSpace: "pre-line" }}>
                      {line}
                    </p>
                  )
              )}
            </div>
          )}
        </div>
      </div>
      <Footer />
    </>
  );
};

export default Subject_detail;
