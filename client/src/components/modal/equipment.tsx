import { DeviceSchema } from '../../schemas/device';
import React, { useEffect } from 'react';

/**
 * @interface Data required for the device information modal.
 * 
 * @param {boolean} showModal if the modal is visible.
 * @param {DeviceSchema} device OLT device information.
 * @param {() => void} onClick Callback to handle the click on the accept button.
 */
interface ModalProps {
    showModal: boolean;
    device: DeviceSchema;
    onClick?: () => void;
}

export default function DeviceInfoModalComponent(content: ModalProps) {

    const handlerAccept = () => {
        document.getElementById('modal-state')?.classList.add('hidden');
        if (content.onClick) content.onClick();
    }

    useEffect(() => {
        const modalAccept = document.getElementById('modal-accept');
        const modalState = document.getElementById('modal-state');

        if (modalAccept && modalState) {
            modalAccept.addEventListener('click', handlerAccept);

            if (content.showModal) modalState.classList.remove('hidden');
            else modalState.classList.add('hidden');
        }

    }, [content.showModal]);

    return(
        <div id="modal-state" className="absolute z-10" aria-labelledby="modal-title" role="dialog" aria-modal="true">
            <aside className="fixed inset-0 bg-black bg-opacity-55 transition-opacity" aria-hidden="true"></aside>
            <div className="fixed inset-0 z-10 w-screen overflow-y-auto">
                <div id="modal-panel" className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
                    <div className="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg">
                        <section className="bg-gray-50 px-4 pb-4 pt-5 sm:p-6 sm:pb-4">
                            <div className="sm:flex sm:items-start">
                                <div className="mt-3 text-center sm:ml-4 sm:mt-0 sm:text-left">
                                    <h3 className="text-2xl font-semibold leading-6 text-blue-700" id="modal-title">{content.device.sysname}</h3>
                                    <div className="mt-4">
                                        {content.device.ip && <p className="text-base text-blue-800 font-bold">IP:
                                            <span className='text-gray-800'> {content.device.ip}</span>
                                        </p>}
                                        {content.device.community && <p className="text-base text-blue-800 font-bold">Comunidad:
                                            <span className='text-gray-800'> {content.device.community}</span>
                                        </p>}
                                        <p className="text-base text-blue-800 font-bold">Estatus:
                                            <span className={`${content.device.is_alive ? "text-green-600" : "text-red-600"}`}> {content.device.is_alive ? "Activo" : "Inactivo"}</span>
                                        </p>
                                        {content.device.syslocation && <p className="text-base text-blue-800 font-bold">Localidad:
                                            <span className='text-gray-800'> {content.device.syslocation}</span>
                                        </p>}
                                        {content.device.last_check && <p className="mt-2 text-xs text-gray-400 font-semibold">Última Revisión del Equipo: {content.device.last_check.toString().split("T")[0]} {content.device.last_check.toString().split("T")[1].split(".")[0]}</p>}
                                    </div>
                                </div>
                            </div>
                        </section>
                        <section className="bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6">
                            <button id="modal-accept" type="button" className="inline-flex w-full justify-center rounded-md bg-blue-600 px-3 py-2 text-sm font-semibold text-white shadow-sm transition-all ease-linear duration-200 hover:bg-blue-700 sm:ml-3 sm:w-auto">Cerrar</button>
                        </section>
                    </div>
                </div>
            </div>
        </div>
    );
}