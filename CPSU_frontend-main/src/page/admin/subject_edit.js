import React, { useState, useEffect } from "react";
import axios from "axios";
import { useParams, useNavigate } from "react-router-dom";
import Headers from "../../component/header";
import Navbar from "../../component/navbar";
import Footer from "../../component/footer";
import Menu from "../../component/menu";
import "../../css/admin/personnel.css";

const Subject_edit = () => {
  const { id } = useParams();
  const navigate = useNavigate();

  const [formData, setFormData] = useState({
    subject_id: "",
    course_id: "",
    thai_course: "",
    plan_type: "",
    semester: "",
    thai_subject: "",
    eng_subject: "",
    credits: "",
    compulsory_subject: "",
    condition: "",
    description_id: "",
    description_thai: "",
    description_eng: "",
    clo_id: "",
    clo: "",
  });

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchSubject = async () => {
      try {
        const res = await axios.get(`http://localhost:8080/api/v1/admin/subject/${id}`);
        setFormData({
          subject_id: res.data.subject_id || "",
          course_id: res.data.course_id || "",
          thai_course: res.data.thai_course || "",
          plan_type: res.data.plan_type || "",
          semester: res.data.semester || "",
          thai_subject: res.data.thai_subject || "",
          eng_subject: res.data.eng_subject || "",
          credits: res.data.credits || "",
          compulsory_subject: res.data.compulsory_subject || "",
          condition: res.data.condition || "",
          description_id: res.data.description_id || "",
          description_thai: res.data.description_thai || "",
          description_eng: res.data.description_eng || "",
          clo_id: res.data.clo_id || "",
          clo: res.data.clo || "",
        });
      } catch (err) {
        console.error("Error loading subject data", err);
        setError("ไม่พบข้อมูลรายวิชาหรือเกิดข้อผิดพลาดในการโหลดข้อมูล");
      } finally {
        setLoading(false);
      }
    };

    fetchSubject();
  }, [id]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSave = async (e) => {
    e.preventDefault();
    setError("");

    try {
      await axios.put(`http://localhost:8080/api/v1/admin/subject/${id}`, formData);
      alert("แก้ไขข้อมูลรายวิชาสำเร็จ");
      navigate("/admin/subject");
    } catch (err) {
      console.error(err);
      setError("เกิดข้อผิดพลาดในการบันทึกข้อมูล");
    }
  };

  const handleCancel = () => {
    navigate("/admin/subject");
  };

  if (loading) return <p>กำลังโหลดข้อมูล...</p>;

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
            <h2>แก้ไขข้อมูลรายวิชา</h2>
            {error && <div className="alert alert-danger">{error}</div>}
            <form onSubmit={handleSave}>
              {[
                { label: "รหัสวิชา", name: "subject_id" },
                { label: "รหัสหลักสูตร", name: "course_id" },
                { label: "ชื่อหลักสูตร (ไทย)", name: "thai_course" },
                { label: "แผนการเรียน", name: "plan_type" },
                { label: "ภาคเรียน", name: "semester" },
                { label: "ชื่อวิชา (ไทย)", name: "thai_subject" },
                { label: "ชื่อวิชา (อังกฤษ)", name: "eng_subject" },
                { label: "หน่วยกิต", name: "credits" },
                { label: "บังคับ/ไม่บังคับ", name: "compulsory_subject" },
                { label: "เงื่อนไขรายวิชา", name: "condition" },
                { label: "รหัสคำอธิบายรายวิชา", name: "description_id" },
                { label: "คำอธิบายรายวิชา (ไทย)", name: "description_thai", isTextArea: true },
                { label: "คำอธิบายรายวิชา (อังกฤษ)", name: "description_eng", isTextArea: true },
                { label: "รหัส CLO", name: "clo_id" },
                { label: "CLO (ผลลัพธ์การเรียนรู้)", name: "clo", isTextArea: true },
              ].map((field) => (
                <div className="mb-3" key={field.name}>
                  <label>{field.label}</label>
                  {field.isTextArea ? (
                    <textarea
                      name={field.name}
                      value={formData[field.name]}
                      onChange={handleChange}
                      className="form-control"
                      rows="4"
                    />
                  ) : (
                    <input
                      type="text"
                      name={field.name}
                      value={formData[field.name]}
                      onChange={handleChange}
                      className="form-control"
                    />
                  )}
                </div>
              ))}

              <button type="submit" className="btn btn-success me-2">
                บันทึก
              </button>
              <button type="button" className="btn btn-secondary" onClick={handleCancel}>
                ยกเลิก
              </button>
            </form>
          </div>
        </div>
      </div>
      <Footer />
    </>
  );
};

export default Subject_edit;
