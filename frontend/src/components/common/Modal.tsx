import React, { Fragment, useState } from "react";

import "../../styles/common/modal_style.css";

interface Props {
  isOpen: boolean;
  onClose: () => void;
  children: React.ReactNode;
  header_title: string;
}

const Modal = ({ isOpen, onClose, children, header_title }: Props) => {
  if (!isOpen) {
    return null;
  }

  return (
    <Fragment>
        <div className="modal_wrap">
            <div className="modal_header">
              <div className="modal_header_title">
                { header_title }
              </div>
              <span className="modal_close_button" onClick={onClose}>
                  &times;
              </span>
            </div>
            <div className="modal_content">
                {children}
            </div>
        </div>
        <div className={`modal_overlay ${isOpen == true ? 'active' : ''}`} onClick={onClose}></div>
    </Fragment>
  );
};

export default Modal;