import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

function EventFormAdd() {
  const [title, setTitle] = useState("");
  const [startDate, setStartDate] = useState("");
  const [endDate, setEndDate] = useState("");
  const [description, setDescription] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    // ตรวจสอบข้อมูล
    if (!title || !startDate || !endDate) {
      alert("กรุณากรอกข้อมูลให้ครบถ้วน");
      return;
    }

    if (new Date(startDate) > new Date(endDate)) {
      alert("วันเริ่มต้นต้องไม่มากกว่าวันสิ้นสุด");
      return;
    }

    const toISODate = (dateStr) => {
      return new Date(dateStr).toISOString(); // "2025-10-21T00:00:00.000Z"
    };

    const newEvent = {
      title,
      detail: description,
      start_date: toISODate(startDate),
      end_date: toISODate(endDate),
    };
    console.log(newEvent);
    try {
      await axios.post("http://localhost:8080/api/v1/admin/calendar", newEvent);
      alert("เพิ่มกิจกรรมสำเร็จ");
      navigate("/admin/calendar");
    } catch (error) {
      console.error("เกิดข้อผิดพลาดในการเพิ่มกิจกรรม:", error);
      alert("ไม่สามารถเพิ่มกิจกรรมได้ กรุณาลองใหม่อีกครั้ง");
    }
  };

  return (
    <div className="container mt-4">
      <h3>เพิ่มกิจกรรม</h3>
      <form onSubmit={handleSubmit}>
        <div className="form-group mb-3">
          <label>ชื่อกิจกรรม</label>
          <input
            type="text"
            className="form-control"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
          />
        </div>

        <div className="form-group mb-3">
          <label>วันเริ่มต้น</label>
          <input
            type="date"
            className="form-control"
            value={startDate}
            onChange={(e) => setStartDate(e.target.value)}
            required
          />
        </div>

        <div className="form-group mb-3">
          <label>วันสิ้นสุด</label>
          <input
            type="date"
            className="form-control"
            value={endDate}
            onChange={(e) => setEndDate(e.target.value)}
            required
          />
        </div>

        <div className="form-group mb-3">
          <label>รายละเอียดกิจกรรม</label>
          <textarea
            className="form-control"
            rows="4"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
          />
        </div>

        <button type="submit" className="btn btn-primary">
          บันทึกกิจกรรม
        </button>
      </form>
    </div>
  );
}

export default EventFormAdd;
