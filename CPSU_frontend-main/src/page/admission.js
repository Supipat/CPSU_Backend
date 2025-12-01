import React from "react";
import Navbar from "../component/navbar";
import Headers from "../component/header";
import Footer from "../component/footer";

const Admission = () => {
  return (
    <>
      <Headers />
      <Navbar />

      <div className="container py-5">
        <h2 className="text-center mb-4">ข้อมูลการรับสมัคร</h2>

        <div className="accordion" id="admissionAccordion">
          {/* รอบที่ 1 */}
          <div className="accordion-item">
            <h2 className="accordion-header" id="headingOne">
              <button
                className="accordion-button"
                type="button"
                data-bs-toggle="collapse"
                data-bs-target="#collapseOne"
                aria-expanded="true"
                aria-controls="collapseOne"
              >
                รอบที่ 1 แฟ้มสะสมผลงาน
              </button>
            </h2>
            <div
              id="collapseOne"
              className="accordion-collapse collapse show"
              aria-labelledby="headingOne"
              data-bs-parent="#admissionAccordion"
            >
              <div className="accordion-body text-start">
                {/* สาขาเทคโนโลยีสารสนเทศ */}
                <h5>สาขาวิชาเทคโนโลยีสารสนเทศ</h5>
                <ul>
                  <li><strong>ภาคปกติ</strong> รับจำนวนตามประกาศ</li>
                  <li><strong>ภาคพิเศษ</strong> รับจำนวนตามประกาศ</li>
                </ul>

                {/* สาขาวิทยาการคอมพิวเตอร์ */}
                <h5 className="mt-3">สาขาวิชาวิทยาการคอมพิวเตอร์</h5>
                <ul>
                  <li><strong>ภาคปกติ</strong> รับจำนวนตามประกาศ</li>
                </ul>

                <h6 className="mt-4">คุณสมบัติ</h6>
                <ul>
                  <li>✅ รับผู้สมัครที่จบจาก รร. หลักสูตรแกนกลาง</li>
                  <li>❌ ไม่รับผู้สมัครจาก รร. หลักสูตรนานาชาติ</li>
                  <li>✅ รับจาก รร. อาชีวะ</li>
                  <li>❌ ไม่รับจาก กศน.</li>
                  <li>❌ ไม่รับ GED</li>
                </ul>

                <h6>เงื่อนไขการรับ</h6>
                <p>
                  ปราศจากโรคหรืออาการของโรค หรือความพิการอันเป็นอุปสรรคต่อการศึกษาและปฏิบัติงานในวิชาชีพนั้น ๆ
                </p>

                <p>
                  🔗 รายละเอียดเพิ่มเติม:{" "}
                  <a
                    href="https://admission.su.ac.th"
                    target="_blank"
                    rel="noopener noreferrer"
                  >
                    https://admission.su.ac.th
                  </a>
                </p>
              </div>
            </div>
          </div>

          {/* รอบที่ 2 */}
          <div className="accordion-item">
            <h2 className="accordion-header" id="headingTwo">
              <button
                className="accordion-button collapsed"
                type="button"
                data-bs-toggle="collapse"
                data-bs-target="#collapseTwo"
                aria-expanded="false"
                aria-controls="collapseTwo"
              >
                รอบที่ 2 โควต้า
              </button>
            </h2>
            <div
              id="collapseTwo"
              className="accordion-collapse collapse"
              aria-labelledby="headingTwo"
              data-bs-parent="#admissionAccordion"
            >
              <div className="accordion-body text-start">
                <h5>สาขาวิชาวิทยาการคอมพิวเตอร์</h5>
                <ul>
                  <li><strong>ภาคพิเศษ:</strong> รับตามประกาศ</li>
                </ul>
              </div>
            </div>
          </div>

          {/* รอบที่ 3 */}
          <div className="accordion-item">
            <h2 className="accordion-header" id="headingThree">
              <button
                className="accordion-button collapsed"
                type="button"
                data-bs-toggle="collapse"
                data-bs-target="#collapseThree"
                aria-expanded="false"
                aria-controls="collapseThree"
              >
                รอบที่ 3 รับตรง
              </button>
            </h2>
            <div
              id="collapseThree"
              className="accordion-collapse collapse"
              aria-labelledby="headingThree"
              data-bs-parent="#admissionAccordion"
            >
              <div className="accordion-body text-start">
                <h5>สาขาวิชาวิทยาการคอมพิวเตอร์</h5>
                <ul>
                  <li><strong>ภาคพิเศษ:</strong> รับตามประกาศ</li>
                </ul>
              </div>
            </div>
          </div>

          {/* รอบที่ 4 */}
          <div className="accordion-item">
            <h2 className="accordion-header" id="headingFour">
              <button
                className="accordion-button collapsed"
                type="button"
                data-bs-toggle="collapse"
                data-bs-target="#collapseFour"
                aria-expanded="false"
                aria-controls="collapseFour"
              >
                รอบที่ 4 รับตรงอิสระ
              </button>
            </h2>
            <div
              id="collapseFour"
              className="accordion-collapse collapse"
              aria-labelledby="headingFour"
              data-bs-parent="#admissionAccordion"
            >
              <div className="accordion-body text-start">
                <h5>สาขาวิชาวิทยาการคอมพิวเตอร์</h5>
                <ul>
                  <li><strong>ภาคพิเศษ:</strong> รับตามประกาศ</li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      </div>

      <Footer />
    </>
  );
};

export default Admission;
