import React, { useEffect, useState } from "react";
import FullCalendar from "@fullcalendar/react";
import dayGridPlugin from "@fullcalendar/daygrid";
import interactionPlugin from "@fullcalendar/interaction";
import axios from "axios";

import "../css/admin/event.css";
import Headers from "../component/header";
import Navbar from "../component/navbar";
import Footer from "../component/footer";

function EventCalendarPublic() {
  const [events, setEvents] = useState([]);
  const [selectedEvent, setSelectedEvent] = useState(null);
  const [showModal, setShowModal] = useState(false);

  // ✅ โหลดกิจกรรมทั้งหมดจาก API
  useEffect(() => {
    const fetchEvents = async () => {
      try {
        const res = await axios.get(
          "http://localhost:8080/api/v1/admin/calendar"
        );

        const mappedEvents = res.data.map((event) => {
          const startDate = new Date(event.start_date);
          const endDate = new Date(event.end_date);

          const isSameDay =
            startDate.toISOString().slice(0, 10) ===
            endDate.toISOString().slice(0, 10);

          return {
            id: event.calendar_id.toString(),
            title: event.title,
            start: startDate.toISOString(),
            ...(isSameDay
              ? { allDay: true } 
              : { end: endDate.toISOString(), allDay: true }),
          };
        });

        setEvents(mappedEvents);
      } catch (error) {
        console.error("ไม่สามารถโหลดข้อมูลกิจกรรม:", error);
      }
    };

    fetchEvents();
  }, []);

  // ✅ เมื่อคลิกกิจกรรม ดึงรายละเอียดจาก API
  const handleEventClick = async (info) => {
    try {
      const res = await axios.get(
        `http://localhost:8080/api/v1/admin/calendar/${info.event.id}`
      );
      const eventData = res.data;
      setSelectedEvent({
        title: eventData.title,
        start: new Date(eventData.start_date).toLocaleString("th-TH"),
        end: new Date(eventData.end_date).toLocaleString("th-TH"),
        description: eventData.detail,
      });
      setShowModal(true);
    } catch (error) {
      console.error("ไม่สามารถโหลดรายละเอียดกิจกรรม:", error);
    }
  };

  const closeModal = () => {
    setShowModal(false);
    setSelectedEvent(null);
  };

  return (
    <>
      <Headers />
      <Navbar />
      <div className="container mt-4">
        <h3 className="mb-4">ปฏิทินกิจกรรม</h3>

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

      {/* ✅ Modal แสดงรายละเอียด */}
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
                <h5 className="modal-title">รายละเอียดกิจกรรม</h5>
                <button
                  type="button"
                  className="btn-close"
                  onClick={closeModal}
                ></button>
              </div>
              <div className="modal-body">
                <p>
                  <strong>ชื่อกิจกรรม:</strong> {selectedEvent.title}
                </p>
                <p>
                  <strong>วันเริ่ม:</strong> {selectedEvent.start}
                </p>
                <p>
                  <strong>วันสิ้นสุด:</strong> {selectedEvent.end}
                </p>
                <p>
                  <strong>รายละเอียด:</strong>
                  <br />
                  {selectedEvent.description || "-"}
                </p>
              </div>
              <div className="modal-footer">
                <button className="btn btn-secondary" onClick={closeModal}>
                  ปิด
                </button>
              </div>
            </div>
          </div>
        </div>
      )}

      <Footer />
    </>
  );
}

export default EventCalendarPublic;
