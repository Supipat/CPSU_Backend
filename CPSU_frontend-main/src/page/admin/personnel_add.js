import React, { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import Headers from "../../component/header";
import Navbar from "../../component/navbar";
import Footer from "../../component/footer";
import Menu from "../../component/menu";
import "../../css/admin/personnel.css";

function PersonnelAdd() {
  const navigate = useNavigate();

  const [formData, setFormData] = useState({
    type_personnel: "สายวิชาการ",
    department_position_name: "",
    thai_academic_position: "",
    eng_academic_position: "",
    thai_name: "",
    eng_name: "",
    education: "",
    related_fields: "",
    email: "",
    website: "",
    file_image: null,
  });

  const [previewImage, setPreviewImage] = useState(null);
  const [error, setError] = useState("");

  const handleChange = (e) => {
    const { name, value, files } = e.target;

    if (name === "file_image") {
      const file = files[0];
      setFormData({ ...formData, file_image: file });
      setPreviewImage(URL.createObjectURL(file));
    } else {
      setFormData({ ...formData, [name]: value });
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const data = new FormData();
    for (let key in formData) {
      data.append(key, formData[key]);
    }

    try {
      await axios.post("http://localhost:8080/api/v1/admin/personnel", data, {
        headers: { "Content-Type": "multipart/form-data" },
      });
      navigate("/admin/personnel");
    } catch (err) {
      console.error("Error adding personnel:", err);
      setError("เกิดข้อผิดพลาดในการเพิ่มบุคลากร");
    }
  };

  return (
    <>
      <Headers />
      <Navbar />

      <div className="container ">
        <div className="row">
          <div className="col-sm-3">
            <Menu />
          </div>

          <div className="col-sm-9">
            <div className="card shadow-sm">
              <div className="card-body">
                <h4 className="card-title mb-4">เพิ่มบุคลากร</h4>

                {error && <div className="alert alert-danger">{error}</div>}

                <form onSubmit={handleSubmit} encType="multipart/form-data">
                  <div className="row">
                    <div className="col-md-6 mb-3">
                      <label>ประเภทบุคลากร</label>
                      <select
                        className="form-select"
                        name="type_personnel"
                        value={formData.type_personnel}
                        onChange={handleChange}
                        required
                      >
                        <option value="สายวิชาการ">สายวิชาการ</option>
                        <option value="สายสนับสนุนวิชาการ">สายสนับสนุนวิชาการ</option>
                      </select>
                    </div>

                    <div className="col-md-6 mb-3">
                      <label>ตำแหน่งในภาควิชา</label>
                      <input
                        type="text"
                        name="department_position_name"
                        className="form-control"
                        value={formData.department_position_name}
                        onChange={handleChange}
                        required
                      />
                    </div>

                    <div className="col-md-6 mb-3">
                      <label>คำนำหน้าทางวิชาการ (TH)</label>
                      <input
                        type="text"
                        name="thai_academic_position"
                        className="form-control"
                        value={formData.thai_academic_position}
                        onChange={handleChange}
                      />
                    </div>

                    <div className="col-md-6 mb-3">
                      <label>Academic Title (EN)</label>
                      <input
                        type="text"
                        name="eng_academic_position"
                        className="form-control"
                        value={formData.eng_academic_position}
                        onChange={handleChange}
                      />
                    </div>

                    <div className="col-md-6 mb-3">
                      <label>ชื่อ-นามสกุล (TH)</label>
                      <input
                        type="text"
                        name="thai_name"
                        className="form-control"
                        value={formData.thai_name}
                        onChange={handleChange}
                        required
                      />
                    </div>

                    <div className="col-md-6 mb-3">
                      <label>Name (EN)</label>
                      <input
                        type="text"
                        name="eng_name"
                        className="form-control"
                        value={formData.eng_name}
                        onChange={handleChange}
                        required
                      />
                    </div>

                    <div className="col-12 mb-3">
                      <label>ประวัติการศึกษา</label>
                      <textarea
                        name="education"
                        className="form-control"
                        rows={3}
                        value={formData.education}
                        onChange={handleChange}
                        placeholder={`แยกบรรทัดด้วย \\n เช่น:\nPh.D. (Computer Science)...`}
                      />
                    </div>

                    <div className="col-12 mb-3">
                      <label>สาขาที่เชี่ยวชาญ</label>
                      <textarea
                        name="related_fields"
                        className="form-control"
                        rows={2}
                        value={formData.related_fields}
                        onChange={handleChange}
                      />
                    </div>

                    <div className="col-md-6 mb-3">
                      <label>Email</label>
                      <input
                        type="email"
                        name="email"
                        className="form-control"
                        value={formData.email}
                        onChange={handleChange}
                      />
                    </div>

                    <div className="col-md-6 mb-3">
                      <label>เว็บไซต์ส่วนตัว</label>
                      <input
                        type="text"
                        name="website"
                        className="form-control"
                        value={formData.website}
                        onChange={handleChange}
                      />
                    </div>

                    <div className="col-md-6 mb-3">
                      <label>อัปโหลดรูปภาพ</label>
                      <input
                        type="file"
                        name="file_image"
                        className="form-control"
                        accept="image/*"
                        onChange={handleChange}
                      />
                    </div>

                    <div className="col-md-6 mb-3 text-center">
                      {previewImage && (
                        <img
                          src={previewImage}
                          alt="Preview"
                          className="img-fluid rounded shadow"
                          style={{ maxHeight: "200px" }}
                        />
                      )}
                    </div>
                  </div>

                  <div className="d-flex justify-content-between mt-4">
                    <button type="submit" className="btn btn-success px-4">
                      บันทึก
                    </button>
                    <button
                      type="button"
                      className="btn btn-secondary px-4"
                      onClick={() => navigate("/admin/personnel")}
                    >
                      ยกเลิก
                    </button>
                  </div>
                </form>
              </div>
            </div>
          </div>
        </div>
      </div>

      <Footer />
    </>
  );
}

export default PersonnelAdd;
