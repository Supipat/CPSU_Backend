import React, { useEffect, useState } from "react";
import FullCalendar from "@fullcalendar/react";
import dayGridPlugin from "@fullcalendar/daygrid";
import interactionPlugin from "@fullcalendar/interaction";
import axios from "axios";
import "../../css/admin/event.css";
import { Link } from "react-router-dom";

import Headers from "../../component/header";
import Navbar from "../../component/navbar";
import Footer from "../../component/footer";
import Menu from "../../component/menu";

function EventCalendarAdmin() {
  const [events, setEvents] = useState([]);
  const [selectedEvent, setSelectedEvent] = useState(null);
  const [showModal, setShowModal] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [editedEvent, setEditedEvent] = useState({});

  useEffect(() => {
    fetchEvents();
  }, []);

  const fetchEvents = async () => {
    try {
      const res = await axios.get(
        "http://localhost:8080/api/v1/admin/calendar"
      );
      const mapped = res.data.map((event) => ({
        id: event.calendar_id,
        title: event.title,
        start: event.start_date,
        end: event.start_date === event.end_date ? null : event.end_date,
        description: event.detail,
        allDay: true,
      }));
      setEvents(mapped);
    } catch (err) {
      console.error("โหลดกิจกรรมล้มเหลว", err);
    }
  };

  const handleEventClick = (info) => {
    const event = events.find((e) => e.id === parseInt(info.event.id));
    setSelectedEvent(event);
    setEditedEvent({
      ...event,
      start: event.start.split("T")[0],
      end: event.end ? event.end.split("T")[0] : event.start.split("T")[0],
    });
    setShowModal(true);
    setIsEditing(false);
  };

  const closeModal = () => {
    setShowModal(false);
    setSelectedEvent(null);
    setIsEditing(false);
  };

  const handleEditChange = (e) => {
    const { name, value } = e.target;
    setEditedEvent({ ...editedEvent, [name]: value });
  };

  const toISODate = (dateStr) => {
    return new Date(dateStr).toISOString(); // "2025-10-21T00:00:00.000Z"
  };

  const handleSave = async () => {
    try {
      const payload = {
        title: editedEvent.title,
        detail: editedEvent.description,
        start_date: toISODate(editedEvent.start),
        end_date: toISODate(editedEvent.end),
      };

      await axios.put(
        `http://localhost:8080/api/v1/admin/calendar/${editedEvent.id}`,
        payload
      );

      alert("บันทึกกิจกรรมสำเร็จ");
      await fetchEvents();
      closeModal();
    } catch (error) {
      console.error("บันทึกไม่สำเร็จ", error);
      alert("บันทึกไม่สำเร็จ กรุณาลองใหม่");
    }
  };

  const handleDelete = async () => {
    const confirmDelete = window.confirm("คุณต้องการลบกิจกรรมนี้หรือไม่?");
    if (!confirmDelete || !selectedEvent) return;

    try {
      await axios.delete(
        `http://localhost:8080/api/v1/admin/calendar/${selectedEvent.id}`
      );

      alert("ลบกิจกรรมสำเร็จ");
      await fetchEvents();
      closeModal();
    } catch (error) {
      console.error("ลบกิจกรรมล้มเหลว", error);
      alert("ไม่สามารถลบกิจกรรมได้ กรุณาลองใหม่");
    }
  };

  return (
    <>
      <Headers />
      <Navbar />
      <div className="container mt-4">
        <div className="row">
          <div className="col-sm-3">
            <Menu />
          </div>
          <div className="col-sm-9">
            <div className="d-flex justify-content-between align-items-center mb-4">
              <h3 className="event-title">ปฏิทินกิจกรรม</h3>
              <Link to="/admin/calendar/add" className="btn-addnews">
                + เพิ่มกิจกรรม
              </Link>
            </div>

            <div className="calendar-wrapper bg-white p-3 shadow rounded mb-4">
              <FullCalendar
                plugins={[dayGridPlugin, interactionPlugin]}
                initialView="dayGridMonth"
                locale="th"
                events={events}
                eventClick={handleEventClick}
                height="auto"
              />
            </div>
          </div>
        </div>
      </div>

      {showModal && selectedEvent && (
        <div
          className="modal fade show"
          style={{ display: "block", backgroundColor: "rgba(0,0,0,0.5)" }}
          tabIndex="-1"
          role="dialog"
        >
          <div className="modal-dialog modal-dialog-centered" role="document">
            <div className="modal-content">
              <div className="modal-header">
                <h5 className="modal-title">
                  {isEditing ? "แก้ไขกิจกรรม" : "รายละเอียดกิจกรรม"}
                </h5>
                <button
                  type="button"
                  className="btn-close"
                  onClick={closeModal}
                ></button>
              </div>
              <div className="modal-body">
                {isEditing ? (
                  <>
                    <div className="mb-2">
                      <label>ชื่อกิจกรรม</label>
                      <input
                        type="text"
                        className="form-control"
                        name="title"
                        value={editedEvent.title}
                        onChange={handleEditChange}
                      />
                    </div>
                    <div className="mb-2">
                      <label>วันเริ่ม</label>
                      <input
                        type="date"
                        className="form-control"
                        name="start"
                        value={editedEvent.start}
                        onChange={handleEditChange}
                      />
                    </div>
                    <div className="mb-2">
                      <label>วันสิ้นสุด</label>
                      <input
                        type="date"
                        className="form-control"
                        name="end"
                        value={editedEvent.end}
                        onChange={handleEditChange}
                      />
                    </div>
                    <div className="mb-2">
                      <label>รายละเอียด</label>
                      <textarea
                        className="form-control"
                        name="description"
                        value={editedEvent.description || ""}
                        onChange={handleEditChange}
                      />
                    </div>
                  </>
                ) : (
                  <>
                    <p>
                      <strong>ชื่อกิจกรรม:</strong> {selectedEvent.title}
                    </p>
                    <p>
                      <strong>วันเริ่ม:</strong> {selectedEvent.start}
                    </p>
                    <p>
                      <strong>วันสิ้นสุด:</strong>{" "}
                      {selectedEvent.end || selectedEvent.start}
                    </p>
                    <p>
                      <strong>รายละเอียด:</strong>
                      <br />
                      {selectedEvent.description || "-"}
                    </p>
                  </>
                )}
              </div>
              <div className="modal-footer">
                {isEditing ? (
                  <>
                    <button className="btn btn-primary" onClick={handleSave}>
                      บันทึก
                    </button>
                    <button
                      className="btn btn-secondary"
                      onClick={() => setIsEditing(false)}
                    >
                      ยกเลิก
                    </button>
                  </>
                ) : (
                  <>
                    <button
                      className="btn btn-warning"
                      onClick={() => setIsEditing(true)}
                    >
                      แก้ไข
                    </button>
                    <button className="btn btn-danger" onClick={handleDelete}>
                      ลบ
                    </button>
                    <button className="btn btn-secondary" onClick={closeModal}>
                      ปิด
                    </button>
                  </>
                )}
              </div>
            </div>
          </div>
        </div>
      )}

      <Footer />
    </>
  );
}

export default EventCalendarAdmin;
