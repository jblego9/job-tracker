import ApplicationsPageControls from "./applications-page-controls";

export default function ApplicationsControlBar({ pageNumber, onPageNumberChange, pageSize, onPageSizeChange }) {
    return (
        <div className="applications-control-bar">
            <ApplicationsPageControls
                pageNumber={pageNumber}
                onPageNumberChange={onPageNumberChange}
                pageSize={pageSize}
                onPageSizeChange={onPageSizeChange}
            />
        </div>
    );
}
