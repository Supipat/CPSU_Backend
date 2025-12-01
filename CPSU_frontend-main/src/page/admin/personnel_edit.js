import React, { useEffect, useState } from "react";
import { useParams, Link, useNavigate } from "react-router-dom";
import axios from "axios";
import "../../css/admin/news.css";
import Headers from "../../component/header";
import Navbar from "../../component/navbar";
import Footer from "../../component/footer";
import Menu from "../../component/menu";

function Edit_Personnel() {
  const { id } = useParams();
  const navigate = useNavigate();

  const [person, setPerson] = useState({
    thai_name: "",
    thai_academic_position: "",
    eng_name: "",
    eng_academic_position: "",
    email: "",
    website: "",
    department_position_name: "",
    type_personnel: "",
    education: "",
    related_fields: "",
    file_image: "",
  });

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [saving, setSaving] = useState(false);

  useEffect(() => {
    const fetchPerson = async () => {
      try {
        const res = await axios.get(
          `http://localhost:8080/api/v1/admin/personnel/${id}`
        );
        setPerson(res.data);
        setLoading(false);
      } catch (err) {
        console.error("Error fetching personnel detail:", err);
        setError("ไม่สามารถโหลดข้อมูลบุคลากรได้");
        setLoading(false);
      }
    };

    fetchPerson();
  }, [id]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setPerson((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setSaving(true);
    try {
      await axios.put(
        `http://localhost:8080/api/v1/admin/personnel/${id}`,
        person
      );
      alert("บันทึกข้อมูลเรียบร้อย");
      navigate("/admin/personnel");
    } catch (err) {
      console.error("Error updating personnel:", err);
      alert("เกิดข้อผิดพลาดในการบันทึกข้อมูล");
    } finally {
      setSaving(false);
    }
  };

  if (loading) {
    return (
      <>
        <Headers />
        <Navbar />
        <div className="container text-center mt-5">
          <p>กำลังโหลดข้อมูล...</p>
        </div>
        <Footer />
      </>
    );
  }

  if (error) {
    return (
      <>
        <Headers />
        <Navbar />
        <div className="container text-center mt-5 text-danger">
          <p>{error}</p>
          <Link to="/admin/personnel" className="btn btn-secondary mt-3">
            ← กลับไปหน้ารายการบุคลากร
          </Link>
        </div>
        <Footer />
      </>
    );
  }

  return (
    <>
      <Headers />
      <Navbar />
      <div className="container-fluid">
        <div className="row">
          <div className="col-sm-3">
            <Menu />
          </div>
          <div className="col-sm-9">
            <form onSubmit={handleSubmit}>
              <h3 className="mb-4">แก้ไขข้อมูลบุคลากร</h3>

              <div className="mb-3 row align-items-center">
                <label className="col-sm-3 col-form-label text-end">
                  ตำแหน่งทางวิชาการ (ไทย)
                </label>
                <div className="col-sm-6">
                  <input
                    type="text"
                    className="form-control"
                    name="thai_academic_position"
                    value={person.thai_academic_position}
                    onChange={handleChange}
                    placeholder="เช่น ผศ.ดร."
                    required
                  />
                </div>
              </div>

              <div className="mb-3 row align-items-center">
                <label className="col-sm-3 col-form-label text-end">
                  ชื่อ-นามสกุล (ไทย)
                </label>
                <div className="col-sm-6">
                  <input
                    type="text"
                    className="form-control"
                    name="thai_name"
                    value={person.thai_name}
                    onChange={handleChange}
                    placeholder="เช่น สิรภัทร แก้วจำลอง"
                    required
                  />
                </div>
              </div>

              <div className="mb-3 row align-items-center">
                <label className="col-sm-3 col-form-label text-end">
                  ตำแหน่งทางวิชาการ (อังกฤษ)
                </label>
                <div className="col-sm-6">
                  <input
                    type="text"
                    className="form-control"
                    name="eng_academic_position"
                    value={person.eng_academic_position}
                    onChange={handleChange}
                    placeholder="เช่น Asst.Prof.Dr."
                  />
                </div>
              </div>

              <div className="mb-3 row align-items-center">
                <label className="col-sm-3 col-form-label text-end">
                  ชื่อ-นามสกุล (อังกฤษ)
                </label>
                <div className="col-sm-6">
                  <input
                    type="text"
                    className="form-control"
                    name="eng_name"
                    value={person.eng_name}
                    onChange={handleChange}
                    placeholder="เช่น Sirak Kaewjamnong"
                  />
                </div>
              </div>

              <div className="mb-3 row align-items-center">
                <label className="col-sm-3 col-form-label text-end">
                  อีเมล
                </label>
                <div className="col-sm-6">
                  <input
                    type="email"
                    className="form-control"
                    name="email"
                    value={person.email}
                    onChange={handleChange}
                    placeholder="example@domain.com"
                    required
                  />
                </div>
              </div>

              <div className="mb-3 row align-items-center">
                <label className="col-sm-3 col-form-label text-end">
                  เว็บไซต์
                </label>
                <div className="col-sm-6">
                  <input
                    type="url"
                    className="form-control"
                    name="website"
                    value={person.website}
                    onChange={handleChange}
                    placeholder="https://example.com"
                  />
                </div>
              </div>

              <div className="mb-3 row align-items-center">
                <label className="col-sm-3 col-form-label text-end">
                  URL รูปภาพ
                </label>
                <div className="col-sm-6">
                  <input
                    type="url"
                    className="form-control"
                    name="file_image"
                    value={person.file_image}
                    onChange={handleChange}
                    placeholder="https://example.com/image.jpg"
                  />
                  <small className="text-muted">
                    ถ้าไม่ใส่จะแสดงรูปโปรไฟล์เริ่มต้น
                  </small>

                  {person.file_image && (
                    <div className="mt-3 text-center">
                      <img
                        src={person.file_image}
                        alt="Preview"
                        style={{
                          maxWidth: "150px",
                          borderRadius: "8px",
                          border: "1px solid #ddd",
                        }}
                        onError={(e) => (e.target.style.display = "none")}
                      />
                    </div>
                  )}
                </div>
              </div>

              <div className="text-end mt-4">
                <Link
                  to="/admin/personnel"
                  className="btn btn-outline-secondary me-2"
                >
                  ยกเลิก
                </Link>
                <button
                  type="submit"
                  className="btn btn-primary"
                  disabled={saving}
                >
                  {saving ? "กำลังบันทึก..." : "บันทึกข้อมูล"}
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
      <Footer />
    </>
  );
}

export default Edit_Personnel;
